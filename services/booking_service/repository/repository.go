package repository

import (
	"context"

	"github.com/piyushsharma67/events_booking/services/booking_service/database"
	"github.com/piyushsharma67/events_booking/services/booking_service/domain"
)

type SeatsRepository interface {
	GenerateSeats(
		ctx context.Context,
		seat []domain.Seat,
	) ([]domain.Seat, error)
}

type RepositoryStruct struct {
	db database.Database
}

func InitialiseRepository(db database.Database) *RepositoryStruct{
	return &RepositoryStruct{
		db: db,
	}
}


