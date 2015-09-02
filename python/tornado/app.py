import os

import tornado.web
import tornado.ioloop
import tornado.options
import tornado.httpserver

from tornado import gen
from tornado.escape import json_encode


import momoko


dsn = 'dbname=test user=test host=localhost port=5432'


class BaseHandler(tornado.web.RequestHandler):
    @property
    def db(self):
        return self.application.db


class PSQLSelectHandler(BaseHandler):
    @gen.coroutine
    def get(self):
        sql = """SELECT id, uid, bool_value, int_value, float_value, text_value, date_time, date, array_data, json_data FROM models LIMIT %s"""
        cur = yield self.db.execute(sql, (3,))
               
        data = []
        for m in cur.fetchall():
            data.append({
                'id': m[0], 
                'uid': m[1],
                'bool_value': m[2],
                'int_value': m[3],
                'float_value': m[4],
                'text_value': m[5],
                'date_time': m[6].isoformat() if m[6] else None,
                'date': m[7].strftime('%Y/%m/%d') if m[7] else None,
                'array_data': m[8],
                'json_data': m[9],
            }) 

        self.set_header('Content-Type', 'application/json')
        self.write(json_encode(data))
        self.finish()


def main():
    try:
        tornado.options.parse_command_line()
        application = tornado.web.Application([
            (r'/psql/select', PSQLSelectHandler),
        ], debug=False)

        ioloop = tornado.ioloop.IOLoop.instance()

        application.db = momoko.Pool(
            dsn=dsn,
            size=1,
            max_size=3,
            ioloop=ioloop,
            setsession=("SET TIME ZONE UTC",),
            raise_connect_errors=False,
        )

        # this is a one way to run ioloop in sync
        future = application.db.connect()
        ioloop.add_future(future, lambda f: ioloop.stop())
        ioloop.start()

        if application.db.server_version >= 90200:
            future = application.db.register_json()
            # This is the other way to run ioloop in sync
            ioloop.run_sync(lambda: future)

        http_server = tornado.httpserver.HTTPServer(application)
        http_server.listen(8000, 'localhost')
        ioloop.start()

    except KeyboardInterrupt:
        print('Exit')


if __name__ == '__main__':
    main()
