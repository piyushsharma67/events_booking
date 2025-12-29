package models

type CreateEventRequest struct {
	Title       string              `json:"title" binding:"required"`
	Description string              `json:"description"`
	ImageURL    string              `json:"image_url"`
	Location    string              `json:"location" binding:"required"`
	StartTime   string              `json:"start_time" binding:"required"` // RFC3339
	EndTime     string              `json:"end_time" binding:"required"`
	Rows        []SeatingRowRequest `json:"rows" binding:"required,min=1"`
}

type SeatingRowRequest struct {
	RowLabel string `json:"row" binding:"required"`
	Seats    int    `json:"seats" binding:"required,min=1"`
}
