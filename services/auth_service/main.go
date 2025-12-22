package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/piyushsharma67/movie_booking/services/auth_service/databases"
	"github.com/piyushsharma67/movie_booking/services/auth_service/repository"
	"github.com/piyushsharma67/movie_booking/services/auth_service/routes"
	"github.com/piyushsharma67/movie_booking/services/auth_service/service"
	"github.com/streadway/amqp"
)

func main() {

	// 1️⃣ Initialize low-level DB (needs Close)
	pgxpool, queries := databases.InitPostgres()
	defer pgxpool.Close()

	// 2️⃣ Wrap with interface
	db := databases.NewPostgresDB(queries)
	repository := repository.NewUserRepository(db)
	user := os.Getenv("RABBITMQ_USER")
	pass := os.Getenv("RABBITMQ_PASSWORD")
	host := os.Getenv("RABBITMQ_HOST")
	port := os.Getenv("RABBITMQ_PORT")
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", user, pass, host, port))
	notifier, err := service.NewRabbitMQNotifier(conn, "notifications")
	if err != nil {
		log.Fatal(err)
	}

	srv := service.NewAuthService(repository, notifier)
	r := routes.InitRoutes(srv)

	log.Println("Server running on :8001")
	log.Fatal(http.ListenAndServe(":8001", r))
}
