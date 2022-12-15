--
-- PostgreSQL database dump
--

-- Dumped from database version 15.1 (Debian 15.1-1.pgdg110+1)
-- Dumped by pg_dump version 15.1 (Debian 15.1-1.pgdg110+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: event_participants; Type: TABLE; Schema: public; Owner: sa
--

CREATE TABLE public.event_participants (
    id integer NOT NULL,
    event_id integer NOT NULL,
    name character varying NOT NULL,
    email character varying NOT NULL,
    status character varying DEFAULT 'REGISTERED'::character varying NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.event_participants OWNER TO sa;

--
-- Name: event_participants_id_seq; Type: SEQUENCE; Schema: public; Owner: sa
--

CREATE SEQUENCE public.event_participants_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.event_participants_id_seq OWNER TO sa;

--
-- Name: event_participants_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sa
--

ALTER SEQUENCE public.event_participants_id_seq OWNED BY public.event_participants.id;


--
-- Name: event_tags; Type: TABLE; Schema: public; Owner: sa
--

CREATE TABLE public.event_tags (
    event_id integer NOT NULL,
    tag_id integer NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.event_tags OWNER TO sa;

--
-- Name: events; Type: TABLE; Schema: public; Owner: sa
--

CREATE TABLE public.events (
    id integer NOT NULL,
    name character varying NOT NULL,
    start_date timestamp without time zone NOT NULL,
    end_date timestamp without time zone NOT NULL,
    description character varying DEFAULT '-'::character varying,
    location character varying,
    created_by integer NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.events OWNER TO sa;

--
-- Name: events_id_seq; Type: SEQUENCE; Schema: public; Owner: sa
--

CREATE SEQUENCE public.events_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.events_id_seq OWNER TO sa;

--
-- Name: events_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sa
--

ALTER SEQUENCE public.events_id_seq OWNED BY public.events.id;


--
-- Name: tags; Type: TABLE; Schema: public; Owner: sa
--

CREATE TABLE public.tags (
    id integer NOT NULL,
    name character varying NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.tags OWNER TO sa;

--
-- Name: tags_id_seq; Type: SEQUENCE; Schema: public; Owner: sa
--

CREATE SEQUENCE public.tags_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.tags_id_seq OWNER TO sa;

--
-- Name: tags_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sa
--

ALTER SEQUENCE public.tags_id_seq OWNED BY public.tags.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: sa
--

CREATE TABLE public.users (
    id integer NOT NULL,
    name character varying NOT NULL,
    username character varying NOT NULL,
    email character varying NOT NULL,
    secret character varying NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.users OWNER TO sa;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: sa
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO sa;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sa
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: event_participants id; Type: DEFAULT; Schema: public; Owner: sa
--

ALTER TABLE ONLY public.event_participants ALTER COLUMN id SET DEFAULT nextval('public.event_participants_id_seq'::regclass);


--
-- Name: events id; Type: DEFAULT; Schema: public; Owner: sa
--

ALTER TABLE ONLY public.events ALTER COLUMN id SET DEFAULT nextval('public.events_id_seq'::regclass);


--
-- Name: tags id; Type: DEFAULT; Schema: public; Owner: sa
--

ALTER TABLE ONLY public.tags ALTER COLUMN id SET DEFAULT nextval('public.tags_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: sa
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: event_participants; Type: TABLE DATA; Schema: public; Owner: sa
--

COPY public.event_participants (id, event_id, name, email, status, created_at, updated_at, deleted_at) FROM stdin;
\.


--
-- Data for Name: event_tags; Type: TABLE DATA; Schema: public; Owner: sa
--

COPY public.event_tags (event_id, tag_id, created_at) FROM stdin;
\.


--
-- Data for Name: events; Type: TABLE DATA; Schema: public; Owner: sa
--

COPY public.events (id, name, start_date, end_date, description, location, created_by, created_at, updated_at, deleted_at) FROM stdin;
\.


--
-- Data for Name: tags; Type: TABLE DATA; Schema: public; Owner: sa
--

COPY public.tags (id, name, created_at, updated_at, deleted_at) FROM stdin;
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: sa
--

COPY public.users (id, name, username, email, secret, created_at, updated_at, deleted_at) FROM stdin;
1	M Agung Hikmatullah	magunghkm	magunghkm@mail.co	kUBhtNE+dB2cq8A5qWwK+rAlULkJkfsUSiF1kjrK0Emk5/DqwEYeyxeW2jDTRsph8MNY80St9X2BgbtrTjMO/oE09NHXz8SsjXRjtEw7mb2SPsvK0nRzc3IoUqABz4NZdqoBhNwSvwsH4OWMSDgu4QMXYo1BriiwoVvNNcjtTtIqa7nR/fuDe8j8D6JLQH0lntYQtuuwqJkCDTgiCwucC3tR4nlTgAvc+dzuR7Nk4ygDCKo9PIAi90SD1vNdZXYC	2022-12-15 21:57:59.006528	2022-12-15 21:57:59.006528	\N
\.


--
-- Name: event_participants_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sa
--

SELECT pg_catalog.setval('public.event_participants_id_seq', 14, true);


--
-- Name: events_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sa
--

SELECT pg_catalog.setval('public.events_id_seq', 4, true);


--
-- Name: tags_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sa
--

SELECT pg_catalog.setval('public.tags_id_seq', 2, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sa
--

SELECT pg_catalog.setval('public.users_id_seq', 10, true);


--
-- Name: event_participants event_participants_pk; Type: CONSTRAINT; Schema: public; Owner: sa
--

ALTER TABLE ONLY public.event_participants
    ADD CONSTRAINT event_participants_pk PRIMARY KEY (id);


--
-- Name: events events_pk; Type: CONSTRAINT; Schema: public; Owner: sa
--

ALTER TABLE ONLY public.events
    ADD CONSTRAINT events_pk PRIMARY KEY (id);


--
-- Name: tags tags_pk; Type: CONSTRAINT; Schema: public; Owner: sa
--

ALTER TABLE ONLY public.tags
    ADD CONSTRAINT tags_pk PRIMARY KEY (id);


--
-- Name: users users_pk; Type: CONSTRAINT; Schema: public; Owner: sa
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pk PRIMARY KEY (id);


--
-- Name: users users_un; Type: CONSTRAINT; Schema: public; Owner: sa
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_un UNIQUE (username);


--
-- Name: events_deleted_at_idx; Type: INDEX; Schema: public; Owner: sa
--

CREATE INDEX events_deleted_at_idx ON public.events USING btree (deleted_at);


--
-- Name: users_deleted_at_idx; Type: INDEX; Schema: public; Owner: sa
--

CREATE INDEX users_deleted_at_idx ON public.users USING btree (deleted_at);


--
-- Name: event_participants event_participants_fk; Type: FK CONSTRAINT; Schema: public; Owner: sa
--

ALTER TABLE ONLY public.event_participants
    ADD CONSTRAINT event_participants_fk FOREIGN KEY (event_id) REFERENCES public.events(id) ON DELETE CASCADE;


--
-- Name: event_tags event_tags_fk; Type: FK CONSTRAINT; Schema: public; Owner: sa
--

ALTER TABLE ONLY public.event_tags
    ADD CONSTRAINT event_tags_fk FOREIGN KEY (event_id) REFERENCES public.events(id) ON DELETE CASCADE;


--
-- Name: event_tags event_tags_fk_1; Type: FK CONSTRAINT; Schema: public; Owner: sa
--

ALTER TABLE ONLY public.event_tags
    ADD CONSTRAINT event_tags_fk_1 FOREIGN KEY (tag_id) REFERENCES public.tags(id) ON DELETE CASCADE;


--
-- Name: events events_fk; Type: FK CONSTRAINT; Schema: public; Owner: sa
--

ALTER TABLE ONLY public.events
    ADD CONSTRAINT events_fk FOREIGN KEY (created_by) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

