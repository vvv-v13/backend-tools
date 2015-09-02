import os
import momoko

import tornado.web
import tornado.ioloop
from tornado.options import define, options
from tornado.web import RequestHandler, url

import tornado.httpserver

from tornado import gen
from tornado.escape import json_encode


define('port', default=8000, help='run on the given port', type=int)
define('dsn', default='dbname=test user=test host=localhost port=5432', help='Database DSN', type=str)


class BaseHandler(RequestHandler):
    @property
    def db(self):
        return self.application.db


class PSQLSelectHandler(BaseHandler):
    @gen.coroutine
    def get(self):
        sql = """
                SELECT
                        id,
                        uid,
                        bool_value,
                        int_value,
                        float_value,
                        text_value,
                        date_time,
                        date,
                        array_data,
                        json_data
                FROM
                    models
                LIMIT %s
        """

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
    options.parse_command_line()

    application = tornado.web.Application([
        (r'/psql/select', PSQLSelectHandler),
    ], debug=False)

    ioloop = tornado.ioloop.IOLoop.instance()

    application.db = momoko.Pool(
        dsn=options.dsn,
        size=1,
        max_size=3,
        ioloop=ioloop,
    )

    # this is a one way to run ioloop in sync
    future = application.db.connect()
    ioloop.add_future(future, lambda f: ioloop.stop())
    ioloop.start()

    http_server = tornado.httpserver.HTTPServer(application)
    http_server.listen(options.port, 'localhost')
    ioloop.start()


if __name__ == '__main__':
    try:
        main()
    except KeyboardInterrupt:
        print('Exit')

