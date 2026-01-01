package que

type GenerateSeatsMessage struct {
	EventID    string `json:"event_id"`
	SeatLayout struct {
		Rows []struct {
			Row   string `json:"row"`
			Count int    `json:"count"`
		} `json:"rows"`
	} `json:"seat_layout"`
}
