package utils

import (
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func ToText(s string) pgtype.Text {
	if s == "" {
		return pgtype.Text{Valid: false}
	}
	return pgtype.Text{String: s, Valid: true}
}

func GetPgTime(timeStr string) (pgtype.Timestamptz,error) {
	if timeStr == "" {
		return pgtype.Timestamptz{Valid: false}, nil
	}
	t, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		return pgtype.Timestamptz{},fmt.Errorf("invalid start_time format: %w", err)
	}

	startTime := pgtype.Timestamp{
		Time:  t,
		Valid: true,
	}

	return pgtype.Timestamptz(startTime),nil
}
