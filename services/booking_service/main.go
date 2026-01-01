package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/piyushsharma67/events_booking/services/booking_service/database"
	"github.com/piyushsharma67/events_booking/services/booking_service/que"
	"github.com/piyushsharma67/events_booking/services/booking_service/repository"
	"github.com/piyushsharma67/events_booking/services/booking_service/routes"
	"github.com/piyushsharma67/events_booking/services/booking_service/service"
	"github.com/piyushsharma67/events_booking/services/booking_service/utils"
	"github.com/streadway/amqp"
)

func main() {

	r := routes.InitRoutes()
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = utils.DEFAULT_SERVER_PORT

	}
	httpServer := http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	db, err := database.NewPostgres()
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}

	repository := repository.InitialiseRepository(db)
	srv := service.InitialiseService(repository)

	// ---------- RabbitMQ ----------
	rabbitURL := utils.BuildRabbitURL()
	if rabbitURL == "" {
		rabbitURL = "amqp://guest:guest@rabbitmq:5672/"
	}

	conn, err := amqp.Dial(rabbitURL)
	if err != nil {
		log.Fatal("rabbitmq connection failed:", err)
	}

	generateSeatsConsumer, err := que.NewBookingConsumer(conn, os.Getenv("RABBITMQ_QUEUE_GENERATE_SEATS"), que.GenerateSeatsHandler(srv))
	if err != nil {
		log.Fatal("generate seats consumer init failed:", err)
	}

	go func() {
		err := httpServer.ListenAndServe()

		if err != http.ErrServerClosed {
			fmt.Println("http error listen and server", err.Error())
		}
	}()

	// ---------- Context ----------
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		log.Println("RabbitMQ generate seats consumer started")
		if err := generateSeatsConsumer.Start(ctx); err != nil {
			log.Fatal("generate seats consumer error:", err)
		}
	}()

	termSign := make(chan os.Signal, 1)

	signal.Notify(termSign, syscall.SIGINT)
	signal.Notify(termSign, syscall.SIGTERM)

	<-termSign
	cancel()

	if err := httpServer.Shutdown(context.Background()); err != nil {
		log.Println("Server shutdown", err.Error())
	}

	conn.Close()

}
