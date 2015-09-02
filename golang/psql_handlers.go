package main

import (
	"database/sql"
	"encoding/json"
	"github.com/lib/pq"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Model struct {
	Id           int              `json:"id"`
	Uid          string           `json:"uid"`
	DateTime     time.Time        `json:"dateTime"`
	DateTimeText string           `json:"dateTimeText"`
	Date         time.Time        `json:"date"`
	DateText     string           `json:"dateText"`
	Text         string           `json:"text"`
	Int          int              `json:"int"`
	Float        float32          `json:"float"`
	Bool         bool             `json:"bool"`
	JsonData     *json.RawMessage `json:"json"`
	ArrayData    []int            `json:"array"`
}

func strToIntSlice(s string) []int {
	r := strings.Trim(s, "{}")
	a := make([]int, 0)
	for _, t := range strings.Split(r, ",") {
		i, err := strconv.Atoi(t)
		if err == nil {
			a = append(a, i)
		}
	}
	return a
}

// PostgreSQL SELECT handler
func psqlSelectHandler(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "GET" {
			http.Error(w, http.StatusText(405), 405)
			return
		}

		rows, err := db.Query(`SELECT id, uid, bool_value, int_value, float_value, text_value, date_time, date, array_data, json_data FROM models LIMIT $1`, 5)

		if err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		defer rows.Close()

		models := make([]*Model, 0)

		for rows.Next() {
			model := new(Model)
			var uid, text_value, array_data, json_data sql.NullString
			var int_value sql.NullInt64
			var float_value sql.NullFloat64
			var bool_value sql.NullBool
			var date_time, date pq.NullTime

			err := rows.Scan(&model.Id, &uid, &bool_value, &int_value, &float_value, &text_value, &date_time, &date, &array_data, &json_data)
			if err != nil {
				log.Println(err)
				http.Error(w, http.StatusText(500), 500)
				return
			}

			model.Uid = uid.String
			model.Text = text_value.String
			model.Int = int(int_value.Int64)
			model.Float = float32(float_value.Float64)
			model.Bool = bool_value.Bool
			model.DateTime = date_time.Time
			model.Date = date.Time

			if date_time.Valid == true {
				model.DateTime = date_time.Time
				model.DateTimeText = date_time.Time.Format("2006-01-02 15:04:01")
			}

			if date.Valid == true {
				model.DateText = date.Time.Format("2006-01-02")
			}

			if json_data.Valid == true {
				var j *json.RawMessage
				json.Unmarshal([]byte(json_data.String), &j)
				model.JsonData = j
			}

			if array_data.Valid == true {
				model.ArrayData = strToIntSlice(array_data.String)
			}

			models = append(models, model)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(models)
	})
}

// PostgreSQL INSERT Handler
func psqlInsertHandler(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		_, err := db.Exec("INSERT INTO models (uid, data) VALUES($1, $2)", "c8cb5c0b-a42e-4490-8e3d-f617f439dc27", "data")

		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}

		/*
		   fmt.Println("# Deleting")
		    stmt, err = db.Prepare("delete from userinfo where uid=$1")
		    checkErr(err)

		    res, err = stmt.Exec(lastInsertId)
		    checkErr(err)

		    affect, err = res.RowsAffected()
		    checkErr(err)

		    fmt.Println(affect, "rows changed")
		*/

		w.Header().Set("Content-Type", "application/json")
		//json.NewEncoder(w).Encode(users)
	})
}
