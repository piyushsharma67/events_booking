package service

import (
	"context"
	"encoding/json"

	"github.com/piyushsharma67/events_booking/services/events_service/models"
	"github.com/piyushsharma67/events_booking/services/events_service/que"
	"github.com/piyushsharma67/events_booking/services/events_service/repository"
)

type EventService struct {
	Repository repository.EventRepository
	Publisher  *que.QuePublisher
}

func GetEventService(repository repository.EventRepository, publisher *que.QuePublisher) *EventService {
	return &EventService{
		Repository: repository,
		Publisher:  publisher,
	}
}

func (s *EventService) CreateEvent(
	ctx context.Context,
	req *models.CreateEventRequest,
	organiserId string,
) (*models.EventDocument, error) {

	eventDoc, err := models.MapCreateRequestToDocument(req, organiserId)
	if err != nil {
		return nil, err
	}

	created, err := s.Repository.GenerateEvent(ctx, eventDoc)
	if err != nil {
		return nil, err
	}

	return created, nil
}

func (s *EventService) CreateEventAndGenerateSeats(
	ctx context.Context,
	req *models.CreateEventRequest,
	organiserId string,
) (*models.EventDocument, error) {

	// 1️⃣ Create event document
	eventDoc, err := models.MapCreateRequestToDocument(req, organiserId)
	if err != nil {
		return nil, err
	}

	// 2️⃣ Persist event
	created, err := s.Repository.GenerateEvent(ctx, eventDoc)
	if err != nil {
		return nil, err
	}

	// 3️⃣ Marshal event for RabbitMQ
	payload, err := json.Marshal(created)
	if err != nil {
		return created, err
	}

	// 3️⃣ Publish seats generation event
	err = s.Publisher.Publish(ctx,payload)
	if err != nil {
		// IMPORTANT: do NOT rollback DB here
		// Log & retry via DLQ if needed
		return created, err
	}

	return created, nil
}
