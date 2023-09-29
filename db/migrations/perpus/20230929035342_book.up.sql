-- public.book definition

-- Drop table

-- DROP TABLE public.book;

CREATE TABLE public.book (
	id varchar(255) NOT NULL,
	isbn int4 NOT NULL,
	title varchar(100) NOT NULL,
	genre varchar(100) NOT NULL,
	author varchar(100) NOT NULL,
	publisher varchar(100) NOT NULL,
	edition int4 NOT NULL,
	"year" int4 NOT NULL,
	description varchar(1000) NOT NULL,
	created_by varchar(128) NULL,
	updated_by varchar(128) NULL,
	deleted_by varchar(128) NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT book_pkey PRIMARY KEY (id)
);