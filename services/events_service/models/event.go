package models

type Timestamps struct {
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	DeletedAt *string `json:"deleted_at,omitempty"`
}

type SeatingRowRequest struct {
	RowLabel string `json:"row"`
	Seats    int    `json:"seats"`
}

type Event struct {
	ID          string              `json:"id"`
	Title       string              `json:"title" binding:"required"`
	Description string              `json:"description"`
	ImageURL    string              `json:"image_url"`
	Location    string              `json:"location" binding:"required"`
	StartTime   string              `json:"start_time" binding:"required"`
	EndTime     string              `json:"end_time" binding:"required"`
	Rows        []SeatingRowRequest `json:"rows" binding:"required,min=1"`
	Timestamps
}
