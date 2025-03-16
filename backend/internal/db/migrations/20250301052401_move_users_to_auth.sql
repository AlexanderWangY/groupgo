-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA IF NOT EXISTS auth;

ALTER TABLE users SET SCHEMA auth;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE auth.users SET SCHEMA public;
DROP SCHEMA IF EXISTS auth CASCADE;
-- +goose StatementEnd
