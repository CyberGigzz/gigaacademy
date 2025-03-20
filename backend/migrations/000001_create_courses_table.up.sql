BEGIN;

CREATE TABLE IF NOT EXISTS courses (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    code TEXT UNIQUE,
    credits INT,
    start_date TIMESTAMP,
    end_date TIMESTAMP,
    max_capacity INT,
    status TEXT DEFAULT 'active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

COMMIT;