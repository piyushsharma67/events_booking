package repository

import (
	"context"

	"github.com/piyushsharma67/events_booking/services/booking_service/domain"
)

func (r *RepositoryStruct) GenerateSeats(
	ctx context.Context,
	seats []domain.Seat,
) ([]domain.Seat, error) {

	// Call DB layer
	if err := r.db.GenerateSeatsInDB(seats); err != nil {
		return nil, err
	}

	// Return persisted seats
	return seats, nil
}
