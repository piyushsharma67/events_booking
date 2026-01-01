CREATE EXTENSION IF NOT EXISTS "uuid-ossp";  -- enable uuid generation

CREATE TABLE IF NOT EXISTS seats (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    event_id TEXT NOT NULL,
    row_id TEXT NOT NULL,
    seat_number INT NOT NULL,
    status TEXT NOT NULL CHECK (status IN ('AVAILABLE', 'BOOKED', 'LOCKED')),
    UNIQUE (event_id, row_id, seat_number)
);

CREATE TABLE IF NOT EXISTS bookings (
    id UUID PRIMARY KEY,
    event_id TEXT NOT NULL,
    user_id TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT now()
);
