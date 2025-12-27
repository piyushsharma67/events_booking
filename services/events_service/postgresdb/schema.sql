-- name: InitUsersTable :exec
CREATE TABLE IF NOT EXISTS events (
    id SERIAL PRIMARY KEY,
    organizer_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    title TEXT NOT NULL,
    description TEXT,
    image_url TEXT,
    location TEXT NOT NULL,
    start_time TIMESTAMPTZ NOT NULL,
    end_time TIMESTAMPTZ NOT NULL,
    status TEXT NOT NULL DEFAULT 'draft', -- draft | published | cancelled
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);
