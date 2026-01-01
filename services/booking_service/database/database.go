package database

import "github.com/piyushsharma67/events_booking/services/booking_service/domain"

type Database interface{
	GenerateSeatsInDB(seats []domain.Seat) error
}