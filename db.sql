CREATE TABLE models (
    id serial,
    uid uuid,
    int_value integer,
    float_value double precision,
    text_value text,
    json_data jsonb,
    array_data integer[],
    date_time timestamp with time zone,
    date date,
    bool_value boolean
);


COPY models (id, uid, int_value, float_value, text_value, json_data, array_data, date_time, date, bool_value) FROM stdin;
1	c8cb5c0b-a42e-4490-8e3d-f617f439dc27	5	\N	\N	\N	{}	2015-08-31 20:40:15.04813+03	\N	\N
2	\N	\N	\N	bubu	{}	{4,2,6}	\N	\N	t
3	\N	\N	7.44200000000000017	\N	{"aaa": "bbb"}	\N	\N	2015-08-31	f
\.


SELECT pg_catalog.setval('models_id_seq', 3, true);


