--
-- PostgreSQL database dump
--

-- Dumped from database version 10.2
-- Dumped by pg_dump version 10.2

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = public, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: user; Type: TABLE; Schema: public; Owner: ferdinandusrichard
--

CREATE TABLE "user" (
    id integer NOT NULL,
    fullname character varying(128) NOT NULL,
    email character varying(128) NOT NULL,
    password character varying(255) NOT NULL,
    location character varying(128),
    bio character varying(255),
    web character varying(128),
    verification_token character varying(32),
    reset_password_token character varying(32),
    picture bytea
);


ALTER TABLE "user" OWNER TO ferdinandusrichard;

--
-- Name: user_id_seq; Type: SEQUENCE; Schema: public; Owner: ferdinandusrichard
--

CREATE SEQUENCE user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE user_id_seq OWNER TO ferdinandusrichard;

--
-- Name: user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ferdinandusrichard
--

ALTER SEQUENCE user_id_seq OWNED BY "user".id;


--
-- Name: user id; Type: DEFAULT; Schema: public; Owner: ferdinandusrichard
--

ALTER TABLE ONLY "user" ALTER COLUMN id SET DEFAULT nextval('user_id_seq'::regclass);


--
-- Data for Name: user; Type: TABLE DATA; Schema: public; Owner: ferdinandusrichard
--

COPY "user" (id, fullname, email, password, location, bio, web, verification_token, reset_password_token, picture) FROM stdin;
\.


--
-- Name: user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: ferdinandusrichard
--

SELECT pg_catalog.setval('user_id_seq', 1, false);


--
-- Name: user email_unique; Type: CONSTRAINT; Schema: public; Owner: ferdinandusrichard
--

ALTER TABLE ONLY "user"
    ADD CONSTRAINT email_unique UNIQUE (email);


--
-- Name: user user_pkey; Type: CONSTRAINT; Schema: public; Owner: ferdinandusrichard
--

ALTER TABLE ONLY "user"
    ADD CONSTRAINT user_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

