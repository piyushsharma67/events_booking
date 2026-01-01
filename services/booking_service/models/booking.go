package models

import "time"

type BookingDetails struct {
	ID           interface{} `json:"id"`
	Row          string      `json:"row"`
	SeatNumber   int64       `json:"seat_number"`
	EventDetails interface{} `'json:"events_details"`
}

type Booking struct {
	ID        string
	EventID   string
	UserID    string
	SeatIDs   []int
	Status    string
	CreatedAt time.Time
}
