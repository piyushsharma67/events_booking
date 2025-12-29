package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/piyushsharma67/events_booking/services/events_service/database"
	"github.com/piyushsharma67/events_booking/services/events_service/logger"
	"github.com/piyushsharma67/events_booking/services/events_service/repository"
	"github.com/piyushsharma67/events_booking/services/events_service/routes"
	"github.com/piyushsharma67/events_booking/services/events_service/service"
)

func main() {
	
	logger := logger.NewSlogFileLogger("events", "development", "./logs/events/events.log", slog.LevelInfo)
	logger.Info(fmt.Sprintf("events Server running on port :%s", os.Getenv("SERVER_PORT")))

	/* genrating the db connection */
	// 1️⃣ Initialize low-level DB (needs Close)
	// pgxpool, queries := database.InitPostgres()
	// defer pgxpool.Close()

	mongodbClient,close:=database.ConnectMongo()
	defer close()
	
	db:=database.NewMongoDb(mongodbClient)

	// db:=database.NewSqldb(queries)

	repository:=repository.NewRepos(db)

	srv:=service.GetEventService(*repository)

	r:=routes.InitRoutes(srv,logger)
	

	r.Run("0.0.0.0:8003")

}
