CREATE TABLE IF NOT EXISTS seats (
    id UUID PRIMARY KEY,
    event_id UUID NOT NULL,
    row_id TEXT NOT NULL,
    seat_number INT NOT NULL,
    status TEXT NOT NULL CHECK (status IN ('AVAILABLE', 'BOOKED')),
    UNIQUE (event_id, row_id, seat_number)
);

CREATE TABLE IF NOT EXISTS bookings (
    id UUID PRIMARY KEY,
    event_id UUID NOT NULL,
    user_id UUID NOT NULL,
    created_at TIMESTAMP DEFAULT now()
);
