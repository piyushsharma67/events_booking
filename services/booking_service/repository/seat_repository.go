package repository

import (
	"context"

	"github.com/piyushsharma67/events_booking/services/booking_service/domain"
)

func (r *RepositoryStruct) GenerateSeats(
	ctx context.Context,
	seat []domain.Seat,
) ([]domain.Seat, error){
	return []domain.Seat{},nil
}
