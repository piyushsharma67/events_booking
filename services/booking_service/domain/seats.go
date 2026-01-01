package domain

type Seat struct {
	EventID string
	RowID     string
	SeatNumber  int
	Status  string // AVAILABLE, BOOKED , LOCKED
}