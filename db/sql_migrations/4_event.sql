-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE public.events (
	id serial4 NOT NULL,
	"name" varchar NOT NULL,
	start_date timestamp NOT NULL,
	end_date timestamp NOT NULL,
	description varchar NULL DEFAULT '-'::character varying,
	"location" varchar NULL,
	created_by int4 NOT NULL,
	created_at timestamp NOT NULL DEFAULT now(),
	updated_at timestamp NOT NULL DEFAULT now(),
	deleted_at timestamp NULL,
	CONSTRAINT events_pk PRIMARY KEY (id)
);
CREATE INDEX events_deleted_at_idx ON public.events USING btree (deleted_at);


-- public.events foreign keys

ALTER TABLE public.events ADD CONSTRAINT events_fk FOREIGN KEY (created_by) REFERENCES public.users(id) ON DELETE CASCADE;

-- +migrate StatementEnd