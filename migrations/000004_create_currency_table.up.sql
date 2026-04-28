CREATE TABLE IF NOT EXISTS currencies (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    symbol TEXT NOT NULL,
    icon_url TEXT,
    base_unit_factor INTEGER NOT NULL,
    status BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);