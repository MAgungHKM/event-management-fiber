-- +migrate Up
-- +migrate StatementBegin

INSERT INTO public.users (id, name, username, email, secret, created_at, updated_at, deleted_at) VALUES (1, 'M Agung Hikmatullah', 'magunghkm', 'magunghkm@mail.co', 'umOjpwzbrZeMxE0HjYGyGxISUqs=', '2022-12-15 21:57:59.006528', '2022-12-15 21:57:59.006528', NULL);

-- +migrate StatementEnd