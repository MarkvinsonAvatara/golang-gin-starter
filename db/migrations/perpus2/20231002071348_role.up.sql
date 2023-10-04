BEGIN;

CREATE TABLE IF NOT EXISTS perpus2.role (
    id UUID NOT NULL PRIMARY KEY,
    "name" varchar(255) NOT NULL,
    created_by varchar(128) NULL,
    updated_by varchar(128) NULL,
    deleted_by varchar(128) NULL,
    created_at timestamptz NULL,
    updated_at timestamptz NULL,
    deleted_at timestamptz NULL
);

COMMIT;