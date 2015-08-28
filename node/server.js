var express = require('express'),
    models = require('./models'),
    app = express();
    
app.set('port', process.env.PORT || 8000);


app.get('/psql/select', models.psqSelectAll);

app.listen(app.get('port'), function () {
    console.log('Express server listening on port ' + app.get('port'));
});
