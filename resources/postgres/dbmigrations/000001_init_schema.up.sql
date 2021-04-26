BEGIN;

CREATE TYPE enum_account_type AS ENUM (
    'participant',
    'researcher'
);

CREATE TABLE IF NOT EXISTS accounts (
    id BIGSERIAL Primary Key,
    account_id uuid UNIQUE NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    nonce_first_name VARCHAR(255) UNIQUE NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    nonce_last_name VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    nonce_email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255),
    account_type enum_account_type NOT NULL
);

CREATE INDEX ON "accounts" ("account_id");
CREATE INDEX ON "accounts" ("email");

COMMIT;