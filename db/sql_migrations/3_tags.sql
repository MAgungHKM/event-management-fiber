-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE public.tags (
	id serial4 NOT NULL,
	"name" varchar NOT NULL,
	created_at timestamp NOT NULL DEFAULT now(),
	updated_at timestamp NOT NULL DEFAULT now(),
	deleted_at timestamp NULL,
	CONSTRAINT tags_pk PRIMARY KEY (id)
);

-- +migrate StatementEnd