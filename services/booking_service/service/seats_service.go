package service

import (
	"context"

	"github.com/piyushsharma67/events_booking/services/booking_service/domain"
)

func (s *SeatsService) GenerateSeats(
	ctx context.Context,
	seats []domain.Seat,
) error {
	_, err := s.repo.GenerateSeats(ctx, seats)
	return err
}
