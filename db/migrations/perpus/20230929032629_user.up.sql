-- public."user" definition

-- Drop table

-- DROP TABLE public."user";

CREATE TABLE public."user" (
	id varchar(255) NOT NULL,
	"name" varchar(255) NOT NULL,
	email varchar(255) NOT NULL,
	"password" varchar(255) NOT NULL,
	dob date NOT NULL,
	roleid varchar(255) NULL,
	created_by varchar(128) NULL,
	updated_by varchar(128) NULL,
	deleted_by varchar(128) NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT user_email_key UNIQUE (email),
	CONSTRAINT user_pkey PRIMARY KEY (id)
);


-- public."user" foreign keys

ALTER TABLE public."user" ADD CONSTRAINT user_roleid_fkey FOREIGN KEY (roleid) REFERENCES public."role"(id);