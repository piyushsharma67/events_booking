package database

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
	"github.com/piyushsharma67/events_booking/services/booking_service/domain"
)

type PostgresDb struct{
	db *sql.DB
}
func NewPostgres() (*PostgresDb, error) {
	db, err := sql.Open("postgres", buildDSN())
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	if err := initSchema(db); err != nil {
		return nil, err
	}

	return &PostgresDb{db: db}, nil
}

func buildDSN() string {
	return "postgres://" +
		os.Getenv("DB_USER") + ":" +
		os.Getenv("DB_PASSWORD") + "@" +
		os.Getenv("DB_HOST") + ":" +
		os.Getenv("DB_PORT") + "/" +
		os.Getenv("DB_NAME") +
		"?sslmode=disable"
}


func (p *PostgresDb) GenerateSeats(seats []domain.Seat) error {
	// TODO: insert seats into Postgres
	return nil
}
