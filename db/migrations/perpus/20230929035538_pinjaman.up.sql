-- public.pinjaman definition

-- Drop table

-- DROP TABLE public.pinjaman;

CREATE TABLE public.pinjaman (
	id varchar(255) NOT NULL,
	user_id varchar(255) NOT NULL,
	buku_id varchar(255) NOT NULL,
	tanggal_pinjam date NOT NULL,
	tanggal_kembali date NOT NULL,
	status varchar(255) NOT NULL,
	created_by varchar(128) NULL,
	updated_by varchar(128) NULL,
	deleted_by varchar(128) NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT pinjaman_pkey PRIMARY KEY (id)
);


-- public.pinjaman foreign keys

ALTER TABLE public.pinjaman ADD CONSTRAINT pinjaman_buku_id_fkey FOREIGN KEY (buku_id) REFERENCES public.book(id);
ALTER TABLE public.pinjaman ADD CONSTRAINT pinjaman_user_id_fkey FOREIGN KEY (user_id) REFERENCES public."user"(id);