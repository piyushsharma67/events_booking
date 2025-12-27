-- name: CreateEvent :one
INSERT INTO events (
    organizer_id,
    title,
    description,
    location,
    image_url,
    start_time,
    end_time,
    status
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING
    id,
    organizer_id,
    title,
    description,
    location,
    image_url,
    start_time,
    end_time,
    status,
    created_at,
    updated_at;

-- name: GetEventByID :one
SELECT
    id,
    organizer_id,
    title,
    description,
    location,
    image_url,
    start_time,
    end_time,
    status,
    created_at,
    updated_at
FROM events
WHERE id = $1;

-- name: GetEventsByOrganizer :many
SELECT
    id,
    organizer_id,
    title,
    description,
    location,
    image_url,
    start_time,
    end_time,
    status,
    created_at,
    updated_at
FROM events
WHERE organizer_id = $1
ORDER BY created_at DESC;

-- name: GetPublishedEvents :many
SELECT
    id,
    organizer_id,
    title,
    description,
    location,
    image_url,
    start_time,
    end_time,
    status,
    created_at,
    updated_at
FROM events
WHERE status = 'published'
  AND start_time > NOW()
ORDER BY start_time ASC;

-- name: UpdateEventByOrganizer :one
UPDATE events
SET
    title = $3,
    description = $4,
    location = $5,
    image_url = $6,
    start_time = $7,
    end_time = $8,
    status = $9,
    updated_at = NOW()
WHERE id = $1
  AND organizer_id = $2
RETURNING
    id,
    organizer_id,
    title,
    status,
    updated_at;

-- name: UpdateEventByAdmin :one
UPDATE events
SET
    title = $2,
    description = $3,
    location = $4,
    image_url = $5,
    start_time = $6,
    end_time = $7,
    status = $8,
    updated_at = NOW()
WHERE id = $1
RETURNING
    id,
    organizer_id,
    title,
    status,
    updated_at;

-- name: DeleteEventByOrganizer :exec
DELETE FROM events
WHERE id = $1
  AND organizer_id = $2;

-- name: DeleteEventByAdmin :exec
DELETE FROM events
WHERE id = $1;
