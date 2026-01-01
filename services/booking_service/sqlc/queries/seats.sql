-- name: InsertSeat :exec
INSERT INTO seats (id, event_id, row_id, seat_number, status)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (event_id, row_id, seat_number) DO NOTHING
RETURNING id, event_id, row_id, seat_number, status;

-- name: InsertMultipleSeats :exec
INSERT INTO seats (id, event_id, row_id, seat_number, status)
VALUES 
-- sqlc doesn’t support variadic insert, so you call this in a loop
($1, $2, $3, $4, $5)
ON CONFLICT (event_id, row_id, seat_number) DO NOTHING;
