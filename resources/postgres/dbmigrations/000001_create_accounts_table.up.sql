BEGIN;

CREATE TYPE enum_account_type AS ENUM (
    'PARTICIPANT',
    'RESEARCHER'
);

CREATE TABLE IF NOT EXISTS accounts (
    id SERIAL Primary Key,
    account_id uuid UNIQUE NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255),
    account_type enum_account_type NOT NULL
);

CREATE INDEX ON "accounts" ("email");
CREATE INDEX ON "accounts" ("account_id");

COMMIT;