CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,

    user_id INTEGER NOT NULL,
    wallet_id INTEGER NOT NULL,
    currency_id INTEGER NOT NULL,

    reference TEXT NOT NULL,
    amount BIGINT NOT NULL,

    internal BOOLEAN NOT NULL DEFAULT FALSE,

    transaction_beneficiary_details TEXT,

    status TEXT NOT NULL DEFAULT 'pending',
    transaction_type TEXT NOT NULL,

    metadata TEXT,

    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);