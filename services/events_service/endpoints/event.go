package endpoints

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
	"github.com/piyushsharma67/events_booking/services/events_service/models"
	"github.com/piyushsharma67/events_booking/services/events_service/service"
)

func GenerateEvent(srv *service.EventService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		event := request.(*models.CreateEventRequest)
		// Get organizer ID from context
		organizerID, ok := ctx.Value("user_id").(string)
		if !ok || organizerID == "" {
			return nil, errors.New("organizer ID not found in context")
		}

		genEvent, err := srv.CreateEvent(ctx, event,organizerID)

		if err != nil {
			return nil, err
		}

		return genEvent, nil
	}
}
