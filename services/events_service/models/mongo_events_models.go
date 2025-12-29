package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SeatingRow struct {
	RowLabel  string `bson:"row"`
	SeatCount int    `bson:"seat_count"`
}

type EventDocument struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	ImageURL    string             `bson:"image_url"`
	Location    string             `bson:"location"`
	OrganizerID string             `bson:"organizer_id"`
	StartTime time.Time `bson:"start_time"`
	EndTime   time.Time `bson:"end_time"`
	Rows []SeatingRow `bson:"rows"`
	Status string `bson:"status"` // DRAFT | PUBLISHED | CANCELLED
	CreatedAt time.Time  `bson:"created_at"`
	UpdatedAt time.Time  `bson:"updated_at"`
	DeletedAt *time.Time `bson:"deleted_at,omitempty"`
}
