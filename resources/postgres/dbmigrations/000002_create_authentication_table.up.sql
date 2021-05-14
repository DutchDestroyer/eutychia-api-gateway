BEGIN;

CREATE TABLE IF NOT EXISTS ix_authentication (
    id SERIAL Primary Key,
    ix_session_id uuid UNIQUE NOT NULL,
    auth_token VARCHAR(255) NOT NULL,
    auth_token_key bytea NOT NULL,
    refresh_token VARCHAR(255) NOT NULL,
    refresh_token_key bytea NOT NULL,
    account_id uuid REFERENCES accounts NOT NULL
);

CREATE INDEX ON "ix_authentication" ("account_id");

COMMIT;
