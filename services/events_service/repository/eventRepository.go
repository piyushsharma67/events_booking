package repository

import (
	"context"

	"github.com/piyushsharma67/events_booking/services/events_service/database"
	"github.com/piyushsharma67/events_booking/services/events_service/models"
)

type EventRepository struct {
	Repository database.Database
}

func NewRepos(db database.Database)*EventRepository{
	return &EventRepository{
		Repository: db,
	}
}

func (r *EventRepository) GenerateEvent(ctx context.Context, event *models.Event) (models.Event, error) {
	return r.Repository.GenerateEvent(ctx, event)
}

func (r *EventRepository) DeleteEventByOrganiser(ctx context.Context, id any) error {
	return r.Repository.DeleteEvent(ctx, 1)
}

func (r *EventRepository) DeleteEventByAdmin(ctx context.Context, id any) error {
	return r.Repository.DeleteEvent(ctx, 1)
}

func (r *EventRepository) UpdateEventByOrganiser(ctx context.Context, event *models.Event) (models.Event, error) {
	return r.Repository.UpdateEvent(ctx, event)
}

func (r *EventRepository) UpdateEventByAdmin(ctx context.Context, event *models.Event) (models.Event, error) {
	return r.Repository.UpdateEvent(ctx, event)
}

func (r *EventRepository) GetEvent(ctx context.Context, id any) (models.Event, error) {
	return r.Repository.GetEvent(ctx, 1)
}
