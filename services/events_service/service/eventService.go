package service

import "github.com/piyushsharma67/events_booking/services/events_service/repository"

type EventService struct{
	Repository repository.EventRepository
}

func GetEventService(repository repository.EventRepository)*EventService{
	return &EventService{
		Repository: repository,
	}
}