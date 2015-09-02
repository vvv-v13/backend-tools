### PostgreSQL test table

```
test=# \d+ models
                                                        Table "public.models"
   Column    |           Type           |                      Modifiers                      | Storage  | Stats target | Description 
-------------+--------------------------+-----------------------------------------------------+----------+--------------+-------------
 id          | integer                  | not null default nextval('models_id_seq'::regclass) | plain    |              | 
 uid         | uuid                     |                                                     | plain    |              | 
 int_value   | integer                  |                                                     | plain    |              | 
 float_value | double precision         |                                                     | plain    |              | 
 text_value  | text                     |                                                     | extended |              | 
 json_data   | jsonb                    |                                                     | extended |              | 
 array_data  | integer[]                |                                                     | extended |              | 
 date_time   | timestamp with time zone |                                                     | plain    |              | 
 date        | date                     |                                                     | plain    |              | 
 bool_value  | boolean                  |                                                     | plain    |              | 

```

```
test=# select * from models;
 id |                 uid                  | int_value | float_value | text_value |   json_data    | array_data |          date_time           |    date    | bool_value 
----+--------------------------------------+-----------+-------------+------------+----------------+------------+------------------------------+------------+------------
  1 | c8cb5c0b-a42e-4490-8e3d-f617f439dc27 |         5 |             |            |                | {}         | 2015-08-31 20:40:15.04813+03 |            | 
  2 |                                      |           |             | bubu       | {}             | {4,2,6}    |                              |            | t
  3 |                                      |           |       7.442 |            | {"aaa": "bbb"} |            |                              | 2015-08-31 | f
(3 rows)

```

### API
curl -i http://127.0.0.1:8000/psql/select
