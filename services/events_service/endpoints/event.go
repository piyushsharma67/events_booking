package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/piyushsharma67/events_booking/services/events_service/models"
	"github.com/piyushsharma67/events_booking/services/events_service/service"
)

func SignupEndpoint(srv *service.EventService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(*models.Event)

		return nil, nil // your SQLC User struct can be returned directly
	}
}
