package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/piyushsharma67/events_booking/services/events_service/logger"
	"github.com/piyushsharma67/events_booking/services/events_service/routes"
)

func main() {
	r := routes.InitRoutes()

	logger := logger.NewSlogFileLogger("events", "development", "./logs/events/events.log", slog.LevelInfo)
	logger.Info(fmt.Sprintf("events Server running on port :%s", os.Getenv("SERVER_PORT")))
	fmt.Println("events server is running on port", os.Getenv("SERVER_PORT"))
	// http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")), r)
	r.Run("0.0.0.0:8003")

}
