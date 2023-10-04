BEGIN;

CREATE TABLE IF NOT EXISTS perpus2.user (
    id UUID PRIMARY KEY NOT NULL,
	"name" varchar(255) NOT NULL,
	email varchar(255) NOT NULL,
	"password" varchar(255) NOT NULL,
	dob date NOT NULL,
	"role_id" UUID NULL
        CONSTRAINT roles_id_foreign REFERENCES perpus2.role(id) ON UPDATE CASCADE ON DELETE CASCADE,
	created_by varchar(128) NULL,
	updated_by varchar(128) NULL,
	deleted_by varchar(128) NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL
);

COMMIT;