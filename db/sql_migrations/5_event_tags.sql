-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE public.event_tags (
	event_id int4 NOT NULL,
	tag_id int4 NOT NULL,
	created_at timestamp NOT NULL DEFAULT now()
);


-- public.event_tags foreign keys

ALTER TABLE public.event_tags ADD CONSTRAINT event_tags_fk FOREIGN KEY (event_id) REFERENCES public.events(id) ON DELETE CASCADE;
ALTER TABLE public.event_tags ADD CONSTRAINT event_tags_fk_1 FOREIGN KEY (tag_id) REFERENCES public.tags(id) ON DELETE CASCADE;

-- +migrate StatementEnd