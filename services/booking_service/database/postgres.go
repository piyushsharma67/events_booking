package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/piyushsharma67/events_booking/services/booking_service/domain"
	"github.com/piyushsharma67/events_booking/services/booking_service/sqlc/sqlc_gen"
)

type PostgresDb struct {
	db      *sql.DB
	queries *sqlc_gen.Queries
}

func NewPostgres() (*PostgresDb, error) {
	db, err := sql.Open("postgres", buildDSN())
	if err != nil {
		return nil, err
	}

	// Retry ping until DB is ready
	for i := 0; i < 20; i++ {
		if err := db.Ping(); err == nil {
			break
		}
		log.Println("Postgres not ready, retrying...")
		time.Sleep(2 * time.Second)
	}

	if err := initSchema(db); err != nil {
		return nil, err
	}

	return &PostgresDb{db: db, queries: sqlc_gen.New(db)}, nil
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

func (p *PostgresDb) GenerateSeatsInDB(seats []domain.Seat) error {
	ctx := context.Background()
	for _, s := range seats {
		err := p.queries.InsertSeat(ctx, sqlc_gen.InsertSeatParams{
			ID: uuid.New(),
			EventID:    s.EventID,
			RowID:      s.RowID,
			SeatNumber: s.SeatNumber,
			Status:     "AVAILABLE",
		})
		fmt.Println("the data to be pushed inside", sqlc_gen.InsertSeatParams{
			EventID:    s.EventID,
			RowID:      s.RowID,
			SeatNumber: s.SeatNumber,
			Status:     "AVAILABLE",
		})
		if err != nil {
			fmt.Println("error during insertion", err.Error())
			return err
		}
	}
	return nil
}
