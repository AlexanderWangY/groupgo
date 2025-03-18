-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA IF NOT EXISTS auth;

CREATE TYPE auth.payment_plan AS ENUM ('free', 'basic', 'premium');

CREATE TABLE IF NOT EXISTS auth.users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    name TEXT,
    payment_plan auth.payment_plan NOT NULL DEFAULT 'free',
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- Function for updated_at automation
CREATE OR REPLACE FUNCTION update_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER updated_at_users
BEFORE UPDATE ON auth.users
FOR EACH ROW
EXECUTE FUNCTION update_updated_at();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP FUNCTION IF EXISTS update_updated_at();
DROP TRIGGER IF EXISTS updated_at_users ON auth.users;

DROP TABLE IF EXISTS auth.users;

DROP TYPE IF EXISTS auth.payment_plan;

DROP SCHEMA IF EXISTS auth CASCADE;
-- +goose StatementEnd
