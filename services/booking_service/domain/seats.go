package domain

type Seat struct {
	ID string //optional wont be there diring creation
	EventID string
	RowID     string
	SeatNumber  int32
	Status  string // AVAILABLE, BOOKED , LOCKED
}