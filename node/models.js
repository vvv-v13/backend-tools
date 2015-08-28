var db = require('./pghelper');

function escape(s) {
    return s.replace(/[-\/\\^$*+?.()|[\]{}]/g, '\\$&');
}

function psqSelectAll(req, res, next) {
    var sql = "SELECT id, uid, bool_value, int_value, text_value, date_time, date, array_data FROM models LIMIT 5";
    res.contentType('application/json');

    db.query(sql)
        .then(function (models) {
            return res.send(JSON.stringify(models));
        })
        .catch(next);
};

exports.psqSelectAll = psqSelectAll;
