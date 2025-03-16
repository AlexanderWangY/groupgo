-- +goose Up
-- +goose StatementBegin
CREATE TABLE auth.sessions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES auth.users(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE auth.access_tokens (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    session_id UUID REFERENCES auth.sessions(id) ON DELETE CASCADE,
    token TEXT UNIQUE NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    expires_at TIMESTAMPTZ NOT NULL
);

CREATE TABLE auth.refresh_tokens (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    session_id UUID REFERENCES auth.sessions(id) ON DELETE CASCADE,
    token TEXT UNIQUE NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    expires_at TIMESTAMPTZ NOT NULL,
    is_revoked BOOLEAN DEFAULT false
);

-- Indexes to optimize
CREATE INDEX idx_access_tokens_session_id on auth.access_tokens(session_id);
CREATE INDEX idx_access_tokens_expires_at on auth.access_tokens(expires_at);
CREATE INDEX idx_refresh_tokens_session_id on auth.refresh_tokens(session_id);
CREATE INDEX idx_refresh_tokens_expires_at on auth.refresh_tokens(expires_at);
CREATE INDEX idx_refresh_tokens_is_revoked on auth.refresh_tokens(is_revoked);


CREATE TRIGGER trigger_session_updated_at
BEFORE UPDATE ON auth.sessions
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS auth.sessions;
DROP TABLE IF EXISTS auth.access_tokens;
DROP TABLE IF EXISTS auth.refresh_tokens;
DROP TRIGGER IF EXISTS trigger_session_updated_at;

DROP INDEX IF EXISTS idx_access_tokens_session_id;
DROP INDEX IF EXISTS idx_access_tokens_expires_at;
DROP INDEX IF EXISTS idx_refresh_tokens_session_id;
DROP INDEX IF EXISTS idx_refresh_tokens_expires_at;
DROP INDEX IF EXISTS idx_refresh_tokens_is_revoked;


-- +goose StatementEnd
