package service

import (
	"github.com/piyushsharma67/events_booking/services/booking_service/repository"
)

type SeatsService struct {
	repo repository.SeatsRepository
}

func InitialiseService(repo repository.SeatsRepository)*SeatsService{
	return &SeatsService{
		repo: repo,
	}
}
