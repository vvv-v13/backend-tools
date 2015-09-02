var koa = require('koa'), 
    route = require('koa-route'), 
    app = module.exports = koa() 

var pg = require('co-pg')(require('pg'));


var conn_url = 'postgres://test@localhost:5432/test';

function *query(sql, params) {
    try {
        var conn = yield pg.connect_(conn_url);
        var client = conn[0];
        var done = conn[1];

        var result = yield client.query_(sql, params);
        done();  // Release client back to pool

        return result.rows;
    } catch(ex) {
        console.error(ex.toString());
    }
}


function *psqlSelect() {
    var res = yield query('SELECT * FROM models LIMIT 3;', []);
    this.body = res;
}


app.use(route.get('/psql/select', psqlSelect));

if (!module.parent) app.listen(8000);
