-- +goose Up
-- +goose StatementBegin
CREATE TABLE auth.sessions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES auth.users(id) ON DELETE CASCADE,
    token VARCHAR(255) NOT NULL UNIQUE,
    expires_at TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES auth.users(id)
);

CREATE TABLE auth.refresh_tokens (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES auth.users(id) ON DELETE CASCADE,
    session_id UUID NOT NULL REFERENCES auth.sessions(id) ON DELETE CASCADE,
    token VARCHAR(255) NOT NULL UNIQUE,
    used BOOLEAN NOT NULL DEFAULT FALSE,
    expires_at TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES auth.users(id),
    CONSTRAINT fk_session FOREIGN KEY (session_id) REFERENCES auth.sessions(id)
);

CREATE INDEX idx_refresh_tokens_token ON auth.refresh_tokens(token);
CREATE INDEX idx_refresh_tokens_session_id ON auth.refresh_tokens(session_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS auth.sessions;
DROP TABLE IF EXISTS auth.refresh_tokens;
-- +goose StatementEnd
