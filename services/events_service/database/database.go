package database

import (
	"context"

	"github.com/piyushsharma67/events_booking/services/events_service/models"
)

type Database interface {
	GenerateEvent(ctx context.Context,event *models.Event)(models.Event,error)
	DeleteEvent(ctx context.Context,eventId any)(error)
	UpdateEvent(ctx context.Context,event *models.Event)(models.Event,error)
	GetEvent(ctx context.Context,eventId any)(models.Event,error)
}