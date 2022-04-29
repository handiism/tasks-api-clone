--
-- PostgreSQL database dump
--

-- Dumped from database version 14.2
-- Dumped by pg_dump version 14.2

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
-- Name: list; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.list (
    id bigint NOT NULL,
    user_id uuid NOT NULL,
    title character varying(255) NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE public.list OWNER TO postgres;

--
-- Name: list_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.list_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.list_id_seq OWNER TO postgres;

--
-- Name: list_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.list_id_seq OWNED BY public.list.id;


--
-- Name: subtask; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.subtask (
    id bigint NOT NULL,
    task_id bigint NOT NULL,
    name character varying(255) NOT NULL,
    is_done boolean DEFAULT false NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE public.subtask OWNER TO postgres;

--
-- Name: subtask_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.subtask_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.subtask_id_seq OWNER TO postgres;

--
-- Name: subtask_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.subtask_id_seq OWNED BY public.subtask.id;


--
-- Name: task; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.task (
    id bigint NOT NULL,
    list_id bigint NOT NULL,
    name character varying(255) NOT NULL,
    detail character varying(255),
    due_date timestamp with time zone,
    is_done boolean DEFAULT false NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE public.task OWNER TO postgres;

--
-- Name: task_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.task_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.task_id_seq OWNER TO postgres;

--
-- Name: task_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.task_id_seq OWNED BY public.task.id;


--
-- Name: user; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."user" (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name character varying(255) NOT NULL,
    email character varying(255) NOT NULL,
    password character varying(255) NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE public."user" OWNER TO postgres;

--
-- Name: list id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.list ALTER COLUMN id SET DEFAULT nextval('public.list_id_seq'::regclass);


--
-- Name: subtask id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.subtask ALTER COLUMN id SET DEFAULT nextval('public.subtask_id_seq'::regclass);


--
-- Name: task id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.task ALTER COLUMN id SET DEFAULT nextval('public.task_id_seq'::regclass);


--
-- Data for Name: list; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.list (id, user_id, title, created_at, updated_at) FROM stdin;
1	0803497a-71a0-4578-9dfc-7cc4c1c4730f	autem	2022-03-31 23:50:15.192163+07	2022-03-31 23:50:15.192336+07
2	5ce6cc76-a3e5-40f0-a495-91f9d472379f	atque	2022-03-31 23:50:15.194903+07	2022-03-31 23:50:15.195009+07
3	b09534d7-1877-44de-a161-c4a264f84853	et	2022-03-31 23:50:15.195969+07	2022-03-31 23:50:15.195969+07
4	b09534d7-1877-44de-a161-c4a264f84853	ab	2022-03-31 23:50:15.196544+07	2022-03-31 23:50:15.196969+07
5	cd36341c-50b8-454f-8010-0a10edb72881	incidunt	2022-03-31 23:50:15.197897+07	2022-03-31 23:50:15.19806+07
6	96ed2f85-ed07-42e8-b730-2c5d58064ab1	esse	2022-03-31 23:50:15.198898+07	2022-03-31 23:50:15.199018+07
7	b09534d7-1877-44de-a161-c4a264f84853	voluptas	2022-03-31 23:50:15.199915+07	2022-03-31 23:50:15.200052+07
8	829645e3-40cf-4528-aa0e-c99c5ec782b9	rerum	2022-03-31 23:50:15.200839+07	2022-03-31 23:50:15.200912+07
9	451dd618-8c36-4936-a987-17f5c05b62f6	molestiae	2022-03-31 23:50:15.201658+07	2022-03-31 23:50:15.201729+07
10	21b93de0-8be9-450f-9577-7adc65041800	repudiandae	2022-03-31 23:50:15.202551+07	2022-03-31 23:50:15.202623+07
11	9e75487d-cdf2-4d0a-9f87-993fc50ac82c	enim	2022-03-31 23:50:15.203411+07	2022-03-31 23:50:15.203411+07
12	b160e1d4-349f-44c0-a546-ed554cf42e4e	animi	2022-03-31 23:50:15.204234+07	2022-03-31 23:50:15.204234+07
13	9e75487d-cdf2-4d0a-9f87-993fc50ac82c	reiciendis	2022-03-31 23:50:15.204737+07	2022-03-31 23:50:15.205032+07
14	261d7071-6ec0-4c0f-a216-952da54f3cec	iure	2022-03-31 23:50:15.205238+07	2022-03-31 23:50:15.205738+07
15	3e4d63fb-aafe-4b0f-995f-0f3d812dff98	nesciunt	2022-03-31 23:50:15.206249+07	2022-03-31 23:50:15.206249+07
16	451dd618-8c36-4936-a987-17f5c05b62f6	accusamus	2022-03-31 23:50:15.206992+07	2022-03-31 23:50:15.207126+07
17	829645e3-40cf-4528-aa0e-c99c5ec782b9	quis	2022-03-31 23:50:15.207731+07	2022-03-31 23:50:15.207869+07
18	e4b3bf0a-593a-4dbd-ad25-22e088630245	ab	2022-03-31 23:50:15.208498+07	2022-03-31 23:50:15.208498+07
19	5ce6cc76-a3e5-40f0-a495-91f9d472379f	dignissimos	2022-03-31 23:50:15.209001+07	2022-03-31 23:50:15.209001+07
20	50392f30-0351-4ea0-b6fc-b38c0bc29a61	aut	2022-03-31 23:50:15.209684+07	2022-03-31 23:50:15.209684+07
\.


--
-- Data for Name: subtask; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.subtask (id, task_id, name, is_done, created_at, updated_at) FROM stdin;
1	7	doloribus	f	2022-03-31 23:50:15.223137+07	2022-03-31 23:50:15.223137+07
2	2	eos	f	2022-03-31 23:50:15.224138+07	2022-03-31 23:50:15.224138+07
3	7	ea	t	2022-03-31 23:50:15.224788+07	2022-03-31 23:50:15.224788+07
4	14	neque	f	2022-03-31 23:50:15.225638+07	2022-03-31 23:50:15.225638+07
5	20	sunt	f	2022-03-31 23:50:15.226138+07	2022-03-31 23:50:15.226217+07
6	13	ut	f	2022-03-31 23:50:15.226638+07	2022-03-31 23:50:15.226638+07
7	6	repellendus	t	2022-03-31 23:50:15.227138+07	2022-03-31 23:50:15.227138+07
8	8	ex	f	2022-03-31 23:50:15.227843+07	2022-03-31 23:50:15.227906+07
9	9	aperiam	f	2022-03-31 23:50:15.228346+07	2022-03-31 23:50:15.228346+07
10	5	quibusdam	f	2022-03-31 23:50:15.228346+07	2022-03-31 23:50:15.228845+07
11	5	perspiciatis	f	2022-03-31 23:50:15.228845+07	2022-03-31 23:50:15.228845+07
12	13	corrupti	f	2022-03-31 23:50:15.22942+07	2022-03-31 23:50:15.22942+07
13	17	reprehenderit	f	2022-03-31 23:50:15.229844+07	2022-03-31 23:50:15.229844+07
14	1	dicta	f	2022-03-31 23:50:15.230845+07	2022-03-31 23:50:15.230845+07
15	12	est	f	2022-03-31 23:50:15.231346+07	2022-03-31 23:50:15.23142+07
16	1	illum	t	2022-03-31 23:50:15.231846+07	2022-03-31 23:50:15.231846+07
17	5	eveniet	t	2022-03-31 23:50:15.232345+07	2022-03-31 23:50:15.232345+07
18	6	necessitatibus	f	2022-03-31 23:50:15.232345+07	2022-03-31 23:50:15.232845+07
19	2	eligendi	t	2022-03-31 23:50:15.232845+07	2022-03-31 23:50:15.232845+07
20	3	doloribus	f	2022-03-31 23:50:15.233415+07	2022-03-31 23:50:15.233415+07
\.


--
-- Data for Name: task; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.task (id, list_id, name, detail, due_date, is_done, created_at, updated_at) FROM stdin;
1	12	accusamus	Aut voluptatem consequatur perferendis sit accusantium.	\N	f	2022-03-31 23:50:15.210231+07	2022-03-31 23:50:15.210378+07
2	17	excepturi	\N	\N	f	2022-03-31 23:50:15.212286+07	2022-03-31 23:50:15.212383+07
3	2	fuga	\N	2022-04-03 23:50:15.212885+07	f	2022-03-31 23:50:15.212885+07	2022-03-31 23:50:15.212885+07
4	7	cupiditate	Voluptatem accusantium aut sit perferendis consequatur.	2022-04-04 23:50:15.214386+07	t	2022-03-31 23:50:15.213964+07	2022-03-31 23:50:15.214386+07
5	3	maiores	\N	\N	f	2022-03-31 23:50:15.214907+07	2022-03-31 23:50:15.215007+07
6	15	voluptas	\N	2022-04-06 23:50:15.215501+07	t	2022-03-31 23:50:15.215501+07	2022-03-31 23:50:15.215501+07
7	17	suscipit	Perferendis accusantium consequatur sit aut voluptatem.	\N	f	2022-03-31 23:50:15.216071+07	2022-03-31 23:50:15.216159+07
8	14	sunt	\N	\N	f	2022-03-31 23:50:15.216573+07	2022-03-31 23:50:15.216573+07
9	2	mollitia	\N	\N	f	2022-03-31 23:50:15.217073+07	2022-03-31 23:50:15.217175+07
10	4	aut	Perferendis consequatur voluptatem accusantium sit aut.	\N	f	2022-03-31 23:50:15.217594+07	2022-03-31 23:50:15.217671+07
11	19	consequuntur	\N	\N	f	2022-03-31 23:50:15.218096+07	2022-03-31 23:50:15.218162+07
12	8	sed	\N	\N	f	2022-03-31 23:50:15.218596+07	2022-03-31 23:50:15.218672+07
13	13	et	Voluptatem aut accusantium sit consequatur perferendis.	\N	f	2022-03-31 23:50:15.219122+07	2022-03-31 23:50:15.219225+07
14	19	facere	\N	2022-04-14 23:50:15.219636+07	t	2022-03-31 23:50:15.219636+07	2022-03-31 23:50:15.219636+07
15	16	perferendis	\N	\N	f	2022-03-31 23:50:15.220137+07	2022-03-31 23:50:15.220197+07
16	11	tenetur	Voluptatem sit aut accusantium consequatur perferendis.	\N	f	2022-03-31 23:50:15.220638+07	2022-03-31 23:50:15.220638+07
17	11	ut	\N	2022-04-17 23:50:15.220638+07	t	2022-03-31 23:50:15.220638+07	2022-03-31 23:50:15.220638+07
18	14	ut	\N	\N	f	2022-03-31 23:50:15.221137+07	2022-03-31 23:50:15.221637+07
19	5	et	Sit consequatur accusantium aut voluptatem perferendis.	\N	f	2022-03-31 23:50:15.221637+07	2022-03-31 23:50:15.222137+07
20	18	occaecati	\N	2022-04-20 23:50:15.222137+07	t	2022-03-31 23:50:15.222137+07	2022-03-31 23:50:15.222137+07
\.


--
-- Data for Name: user; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."user" (id, name, email, password, created_at, updated_at) FROM stdin;
e4b3bf0a-593a-4dbd-ad25-22e088630245	Ms. Hailie Sanford	rNIHvjF@lmWnhej.com	GbyvoqClvGiMdQqiiEBEMfqOIfkchkoFqHhvWKDimiDQvIHXeB	2022-03-31 23:50:14.739164+07	2022-03-31 23:50:14.739677+07
96ed2f85-ed07-42e8-b730-2c5d58064ab1	Dr. Gwen Kuhic	LkYEsRA@XsWIPXD.net	iovhYdQoGmvsbyPUCjfDbgsFGfiANmfbKJSMsuseVbeJaIcbmd	2022-03-31 23:50:15.168977+07	2022-03-31 23:50:15.169092+07
71f08a33-8836-43ae-b2bd-8a74ed358266	Queen Florida Ledner	oVSUffB@uYXqBvy.ru	cWqTJKWRSlQwsaKkLefiLeKIWToFWmseZDZXmYVyqXrGpAFtFh	2022-03-31 23:50:15.170935+07	2022-03-31 23:50:15.171421+07
50392f30-0351-4ea0-b6fc-b38c0bc29a61	Ms. Dolly Kuvalis	YwshEQw@qVUmEgL.org	jGFahKANylXEyTpuZRclFvFRsiepgvrOcqQhrfbALVRqmGAMEH	2022-03-31 23:50:15.173006+07	2022-03-31 23:50:15.173006+07
1fb9348c-1eb8-48b1-84aa-afcea4fe863b	Mrs. Mariah Fritsch	kphOHqh@wGsoatR.com	SiWeQCMevdmyTfcybossQGkZNpcsZNVMlvwlSlxgIKeIXVaRUT	2022-03-31 23:50:15.174509+07	2022-03-31 23:50:15.174634+07
21b93de0-8be9-450f-9577-7adc65041800	Miss Justina Senger	QdGYkKg@JkjwXyh.info	DLTIYwtaCAAamAVXUlIUipEouaDRdVbIHVUNwQUeKCZSbDpdyk	2022-03-31 23:50:15.175901+07	2022-03-31 23:50:15.176108+07
8496679d-ce13-4074-b4af-f8cc37e1cbd8	Queen Rachelle Renner	bhOVBrT@vqGNyvl.org	eMQCMuqcNmbWdYMZYXfmAXmeLLQTKfxOKyFxUmOaJiiCMtMiDV	2022-03-31 23:50:15.177223+07	2022-03-31 23:50:15.177395+07
b09534d7-1877-44de-a161-c4a264f84853	Miss Jacynthe Dicki	UWLWjNt@EGDfGEG.ru	NDRtKnCqEYbMPhQZbJuiWNuqayrsGDraJmroVgNOfWgCFBSHbg	2022-03-31 23:50:15.178287+07	2022-03-31 23:50:15.178728+07
5ce6cc76-a3e5-40f0-a495-91f9d472379f	Queen Rosanna Kozey	YRZlcLv@vvvRUFW.com	aONLrfCraOVdswWcOgUTetEnydjARZMJJytqvnfyClPdqUxPgL	2022-03-31 23:50:15.179726+07	2022-03-31 23:50:15.179726+07
829645e3-40cf-4528-aa0e-c99c5ec782b9	Ms. Bianka Quitzon	rtqgwmv@xsKwjCA.com	whucOGwAhtynCjhbKRUTgXLhtKYaCPRZlgaZtHAKokpvPXOSZk	2022-03-31 23:50:15.181111+07	2022-03-31 23:50:15.181111+07
cd36341c-50b8-454f-8010-0a10edb72881	Miss Cecilia Cremin	onnvJeI@xtBSOyI.org	raXjvkIXksqavLhpRvZiLaIWmhOBnKntvZQeCtxKYhGtwWZCQW	2022-03-31 23:50:15.182114+07	2022-03-31 23:50:15.182114+07
ebd045f7-c099-4f50-82e9-3973ee7792b3	Lady Magali Hagenes	GUVRLeU@pFrHKVg.net	SlDJfOfcKfRJuhMTvKxtxYWbwAVYDKomyNkOLleiAjYlbAMmBX	2022-03-31 23:50:15.183114+07	2022-03-31 23:50:15.183114+07
0e5ab4ae-02df-4453-8f8d-4a11044e4bbb	Queen Yvette Hauck	njwtvOJ@EfHiRfr.net	BmFGcfulDPxveBHiebRpENLmaxmAWLZJQjZiHyqmliIYrIicWd	2022-03-31 23:50:15.184352+07	2022-03-31 23:50:15.18448+07
9e75487d-cdf2-4d0a-9f87-993fc50ac82c	Princess Kaycee Halvorson	TaCSFwd@IZDuOEh.biz	omtIOfliPAZJlChYPxYCfcZnxRPCOcIURblfwEtIWOgmakcPXJ	2022-03-31 23:50:15.185271+07	2022-03-31 23:50:15.185271+07
b160e1d4-349f-44c0-a546-ed554cf42e4e	Prof. Icie Boyer	nxCjeWO@gItZUIK.com	faSoYSMsOUWiemJIRvuWhcqdstVVhIremddljSIHhTUgDxNyDM	2022-03-31 23:50:15.186274+07	2022-03-31 23:50:15.186504+07
0803497a-71a0-4578-9dfc-7cc4c1c4730f	Princess Claudine Christiansen	epCqvuH@iSeBSoF.biz	KqquOVynYouUEtmPoBxjqutJGuUfsjChUpZbtADVyyuVCVNxaq	2022-03-31 23:50:15.187372+07	2022-03-31 23:50:15.187474+07
f919f340-4313-4716-aee5-532695d2ee24	Queen Brittany Padberg	Vmkubkt@LKtfUGy.net	CMyocRxrVkgjWCxyTNCnCWAdmSWhHytMLIfRWmMKoQadyyijOj	2022-03-31 23:50:15.18838+07	2022-03-31 23:50:15.188673+07
451dd618-8c36-4936-a987-17f5c05b62f6	Lady Ashleigh Bode	SJSqjGO@MZpwVFr.net	YemPYYOeFUxavmYrljFJRATKZeTFJuhpwQUQDjdWpcYrSyuhef	2022-03-31 23:50:15.189634+07	2022-03-31 23:50:15.189634+07
261d7071-6ec0-4c0f-a216-952da54f3cec	Prof. Elna Beahan	lwihHoA@QgkbKsm.net	DdCXTWsWVXbpELOLYUgUEySiFwUcFpcgBNGFJACJjhTGiGwuit	2022-03-31 23:50:15.190442+07	2022-03-31 23:50:15.190523+07
3e4d63fb-aafe-4b0f-995f-0f3d812dff98	Mrs. Cindy Funk	XvqCPCV@AStVkXy.net	GpdSQIqDlfchvmaUVqPCdKiLIqXBKlZDvkkhsnCtDkNLlMjLsX	2022-03-31 23:50:15.191278+07	2022-03-31 23:50:15.191364+07
\.


--
-- Name: list_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.list_id_seq', 20, true);


--
-- Name: subtask_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.subtask_id_seq', 20, true);


--
-- Name: task_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.task_id_seq', 20, true);


--
-- Name: list list_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.list
    ADD CONSTRAINT list_pkey PRIMARY KEY (id);


--
-- Name: subtask subtask_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.subtask
    ADD CONSTRAINT subtask_pkey PRIMARY KEY (id);


--
-- Name: task task_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.task
    ADD CONSTRAINT task_pkey PRIMARY KEY (id);


--
-- Name: user user_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT user_pkey PRIMARY KEY (id);


--
-- Name: idx_user_email; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_user_email ON public."user" USING btree (email);


--
-- Name: task fk_list_tasks; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.task
    ADD CONSTRAINT fk_list_tasks FOREIGN KEY (list_id) REFERENCES public.list(id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- Name: subtask fk_task_subtasks; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.subtask
    ADD CONSTRAINT fk_task_subtasks FOREIGN KEY (task_id) REFERENCES public.task(id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- Name: list fk_user_lists; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.list
    ADD CONSTRAINT fk_user_lists FOREIGN KEY (user_id) REFERENCES public."user"(id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- PostgreSQL database dump complete
--

