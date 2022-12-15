-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE public.event_participants (
	id serial4 NOT NULL,
	event_id int4 NOT NULL,
	"name" varchar NOT NULL,
	email varchar NOT NULL,
	status varchar NOT NULL DEFAULT 'REGISTERED'::character varying,
	created_at timestamp NOT NULL DEFAULT now(),
	updated_at timestamp NOT NULL DEFAULT now(),
	deleted_at timestamp NULL,
	CONSTRAINT event_participants_pk PRIMARY KEY (id)
);


-- public.event_participants foreign keys

ALTER TABLE public.event_participants ADD CONSTRAINT event_participants_fk FOREIGN KEY (event_id) REFERENCES public.events(id) ON DELETE CASCADE;

-- +migrate StatementEnd