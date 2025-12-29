package service

import (
	"context"

	"github.com/piyushsharma67/events_booking/services/events_service/models"
	"github.com/piyushsharma67/events_booking/services/events_service/repository"
)

type EventService struct {
	Repository repository.EventRepository
}

func GetEventService(repository repository.EventRepository) *EventService {
	return &EventService{
		Repository: repository,
	}
}

func (s *EventService) CreateEvent(
	ctx context.Context,
	req *models.CreateEventRequest,
	organiserId string,
) (*models.EventDocument, error) {

	eventDoc, err := models.MapCreateRequestToDocument(req,organiserId)
	if err != nil {
		return nil, err
	}

	created, err := s.Repository.GenerateEvent(ctx, eventDoc)
	if err != nil {
		return nil, err
	}

	return created, nil
}
