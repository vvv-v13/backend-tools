import ujson
from flask import Flask, jsonify, Response
from flask_sqlalchemy import SQLAlchemy
from sqlalchemy.dialects import postgresql



app = Flask(__name__)
app.config['SQLALCHEMY_DATABASE_URI'] = 'postgresql://test@localhost:5432/test?client_encoding=utf8'


db = SQLAlchemy(app)


class TestModel(db.Model):
    __tablename__ = "models"

    id = db.Column(db.Integer, primary_key=True)
    uid = db.Column(postgresql.UUID) 
    int_value = db.Column(db.Integer)
    text_value = db.Column(db.String)
    bool_value = db.Column(db.Boolean)
    float_value = db.Column(db.Float)
    date_time = db.Column(db.DateTime(timezone=True))
    date = db.Column(db.Date)
    json_data = db.Column(postgresql.JSONB)
    array_data = db.Column(postgresql.ARRAY(db.Integer))


    def as_dict(self):
        return {
            'id': self.id,
            'uid': self.uid,
            'int_value': self.int_value,
            'float_value': self.float_value,
            'bool_value': self.bool_value,
            'text_value': self.text_value,
            'date_time': self.date_time,
            'date': self.date.strftime('%Y/%m/%d') if self.date else None,
            'json_data': self.json_data,
            'array_data': self.array_data,
        }

    @classmethod
    def as_dict_all(cls):
        return [c.as_dict() for c in cls.query.limit(3).all()]


@app.route('/psql/select')
def postgresql():
    resp = TestModel.as_dict_all() 

    resp = Response(response=ujson.dumps(resp), status=200, mimetype="application/json")
    return(resp)
    #return jsonify({'result': resp})


if __name__ == '__main__':
    app.run(debug=True, host="127.0.0.1", port=8000)

