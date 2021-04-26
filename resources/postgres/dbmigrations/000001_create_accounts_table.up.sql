BEGIN;

CREATE TYPE enum_account_type AS ENUM (
    'PARTICIPANT',
    'RESEARCHER'
);

CREATE TABLE IF NOT EXISTS accounts (
    account_id uuid Primary Key NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    nonce_first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    nonce_last_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    nonce_email VARCHAR(255) NOT NULL,
    password VARCHAR(255),
    account_type enum_account_type NOT NULL
);

CREATE INDEX ON "accounts" ("email");

COMMIT;