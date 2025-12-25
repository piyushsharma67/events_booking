package main

import (
	"log"
	"net/http"
	"os"

	"github.com/piyushsharma67/movie_booking/services/movies_service/databases"
	"github.com/piyushsharma67/movie_booking/services/movies_service/routes"
)

func main() {

	// 1️⃣ Initialize low-level DB (needs Close)
	pgxpool, queries := databases.InitPostgres()
	defer pgxpool.Close()

	// 2️⃣ Wrap with interface
	db := databases.NewPostgresDB(queries)

	// 3️⃣ Pass interface to routes
	r := routes.InitialiseRoutes(db)

	log.Println("Server running on :",os.Getenv("SERVER_PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), r))
}
