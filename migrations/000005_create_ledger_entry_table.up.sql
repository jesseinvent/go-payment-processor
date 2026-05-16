
BEGIN;

CREATE TYPE ledger_entry_type AS ENUM ('debit', 'credit');

CREATE TABLE IF NOT EXISTS ledger_entries (
    id              BIGSERIAL PRIMARY KEY,
    user_id         INT NOT NULL,
    wallet_id       INT NOT NULL,
    transaction_id  INT NOT NULL,
    currency_id     INT NOT NULL,
    entry_type      ledger_entry_type NOT NULL,
    amount          BIGINT NOT NULL
    balance_before  BIGINT NOT NULL
    balance_after   BIGINT NOT NULL
    created_at      TIMESTAMP,
    updated_at      TIMESTAMP,
    deleted_at      TIMESTAMP,
);

-- Indexes
CREATE INDEX idx_ledger_entries_transaction_id ON ledger_entries(transaction_id);
CREATE INDEX idx_ledger_entries_currency_id    ON ledger_entries(currency_id);
CREATE INDEX idx_ledger_entries_wallet_id    ON ledger_entries(wallet_id);


COMMIT;