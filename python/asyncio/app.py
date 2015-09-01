import ujson
import asyncio

from aiohttp import web
from aiopg.sa import create_engine

import sqlalchemy as sa
from sqlalchemy.dialects import postgresql


metadata = sa.MetaData()

models = sa.Table(
    'models', metadata,

    sa.Column('id', sa.Integer, primary_key=True),
    sa.Column('uid', postgresql.UUID),
    sa.Column('int_value', sa.Integer),
    sa.Column('text_value', sa.String),
    sa.Column('bool_value', sa.Boolean),
    sa.Column('float_value', sa.Float),
    sa.Column('date_time', sa.DateTime(timezone=True)),
    sa.Column('date', sa.Date),
    sa.Column('json_data', postgresql.JSONB),
    sa.Column('array_data', postgresql.ARRAY(sa.Integer)), 
)


def response_json(data):
    return web.Response(body=ujson.dumps(data).encode('utf-8'), content_type='application/json')


@asyncio.coroutine
def db_psql_middleware(app, handler):
    @asyncio.coroutine
    def middleware(request):
        db = app.get('db_psql')
        if not db:
            app['db_psql'] = db = yield from create_engine(app['psql_dsn'], minsize=1, maxsize=5)
        request.app['db_psql'] = db
        return (yield from handler(request))
    return middleware


@asyncio.coroutine
def psql_select(request):

    with (yield from request.app['db_psql']) as conn:
        result = yield from conn.execute(models.select())

    data = []
    for m in result:
        data.append({
            'id': m.id, 
            'uid': m.uid,
            'int_value': m.int_value,
            'float_value': m.float_value,
            'bool_value': m.bool_value,
            'text_value': m.text_value,
            'date_time': m.date_time,
            'date': m.date.strftime('%Y/%m/%d') if m.date else None,
            'json_data': m.json_data,
            'array_data': m.array_data,
        }) 
    
    return response_json(data)


app = web.Application(middlewares=[db_psql_middleware])
app['psql_dsn'] = 'postgres://test@127.0.0.1:5432/test'


app.router.add_route('GET', '/psql/select', psql_select)
