package models

type GenerateSeats struct {
	EventId string `json:"event_id"`
	RowId   string `json:"row_id"`
	SeatId  int32 `json:"seat_id"`
}