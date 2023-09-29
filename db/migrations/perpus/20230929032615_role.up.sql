-- public."role" definition

-- Drop table

-- DROP TABLE public."role";

CREATE TABLE public."role" (
	id varchar(255) NOT NULL,
	"name" varchar(255) NOT NULL,
	created_by varchar(128) NULL,
	updated_by varchar(128) NULL,
	deleted_by varchar(128) NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT role_pkey PRIMARY KEY (id)
);