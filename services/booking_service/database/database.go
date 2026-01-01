package database

import "github.com/piyushsharma67/events_booking/services/booking_service/domain"

type Database interface{
	GenerateSeats(seats []domain.Seat) error
}