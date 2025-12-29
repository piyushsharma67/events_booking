package database

import (
	"context"

	"github.com/piyushsharma67/events_booking/services/events_service/models"
)

type Database interface {
	GenerateEvent(ctx context.Context,event *models.EventDocument)(*models.EventDocument,error)
	DeleteEvent(ctx context.Context,eventId any)(error)
	UpdateEvent(ctx context.Context,event *models.EventDocument)(*models.EventDocument,error)
	GetEvent(ctx context.Context,eventId any)(*models.EventDocument,error)
}