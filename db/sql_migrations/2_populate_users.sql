-- +migrate Up
-- +migrate StatementBegin

INSERT INTO public.users (id, name, username, email, secret, created_at, updated_at, deleted_at) VALUES (1, 'M Agung Hikmatullah', 'magunghkm', 'magunghkm@mail.co', 'kUBhtNE+dB2cq8A5qWwK+rAlULkJkfsUSiF1kjrK0Emk5/DqwEYeyxeW2jDTRsph8MNY80St9X2BgbtrTjMO/oE09NHXz8SsjXRjtEw7mb2SPsvK0nRzc3IoUqABz4NZdqoBhNwSvwsH4OWMSDgu4QMXYo1BriiwoVvNNcjtTtIqa7nR/fuDe8j8D6JLQH0lntYQtuuwqJkCDTgiCwucC3tR4nlTgAvc+dzuR7Nk4ygDCKo9PIAi90SD1vNdZXYC', '2022-12-15 21:57:59.006528', '2022-12-15 21:57:59.006528', NULL);

-- +migrate StatementEnd