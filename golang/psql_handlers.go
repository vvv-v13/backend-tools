package main

import (
	"database/sql"
	"encoding/json"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Model struct {
	Id   int       `json:"id"`
	Uid  string    `json:"uid"`
	Date time.Time `json:"date"`
	Text string    `json:"text"`
}

//ToNullInt64 validates a sql.NullInt64 if incoming string evaluates to an integer, invalidates if it does not
func ToNullInt64(s string) sql.NullInt64 {
	i, err := strconv.Atoi(s)
	return sql.NullInt64{Int64: int64(i), Valid: err == nil}
}

// PostgreSQL SELECT handler
func psqlSelectHandler(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "GET" {
			http.Error(w, http.StatusText(405), 405)
			return
		}

		rows, err := db.Query(`SELECT id, uid, date, text_data FROM models LIMIT $1`, 5)

		if err != nil {
			panic(err)
			//http.Error(w, http.StatusText(500), 500)
			//return
		}

		defer rows.Close()

		models := make([]*Model, 0)

		for rows.Next() {
			model := new(Model)
			var uid, text sql.NullString
                        var date sql.NullInt64
			err := rows.Scan(&model.Id, &uid, &date, &text)
			if err != nil {
				panic(err)
				//http.Error(w, http.StatusText(500), 500)
				//return
			}

			log.Println(date)

			model.Uid = uid.String
			model.Text = text.String

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

		w.Header().Set("Content-Type", "application/json")
		//json.NewEncoder(w).Encode(users)
	})
}
