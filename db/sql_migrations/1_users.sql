-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE public.users (
	id serial4 NOT NULL,
	"name" varchar NOT NULL,
	username varchar NOT NULL,
	email varchar NOT NULL,
	secret varchar NOT NULL,
	created_at timestamp NOT NULL DEFAULT now(),
	updated_at timestamp NOT NULL DEFAULT now(),
	deleted_at timestamp NULL,
	CONSTRAINT users_pk PRIMARY KEY (id),
	CONSTRAINT users_un UNIQUE (username)
);
CREATE INDEX users_deleted_at_idx ON public.users USING btree (deleted_at);

-- +migrate StatementEnd