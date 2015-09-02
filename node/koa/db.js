var co = require('co');
var pg = require('co-pg')(require('pg'));
var thunkify = require('thunkify');
var path = require('path');
var fs = require('co-fs');
var authen = require('./authen');
var belt = require('./belt');

var conn_url = 'postgres://dbusername:dbpassword@abcd.efgh.us-east-1.rds.amazonaws.com:5432/dbname';

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


exports.find_users = function *() {
  return yield query('SELECT * FROM users;', []);
}

