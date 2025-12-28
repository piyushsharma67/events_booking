package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/piyushsharma67/events_booking/services/events_service/models"
	"github.com/piyushsharma67/events_booking/services/events_service/service"
)

func GenerateEvent(srv *service.EventService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		event := request.(*models.Event)

		genEvent,err:=srv.Repository.GenerateEvent(ctx,event)

		if err!=nil{
			return nil,err
		}

		return genEvent,nil
	}
}
