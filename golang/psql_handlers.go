package main

import (
	"database/sql"
	"encoding/json"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"time"
)

type Model struct {
	Id   int       `json:"id"`
	Uid  string    `json:"uid"`
	Date time.Time `json:"date"`
	Text string    `json:"text"`
	Int  int       `json:"int"`
	Float  float32   `json:"float"`
	Bool  bool   `json:"bool"`
}


// PostgreSQL SELECT handler
func psqlSelectHandler(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "GET" {
			http.Error(w, http.StatusText(405), 405)
			return
		}

		rows, err := db.Query(`SELECT id, uid, bool_value, int_value, text_value FROM models LIMIT $1`, 5)

		if err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		defer rows.Close()

		models := make([]*Model, 0)

		for rows.Next() {
			model := new(Model)
			var uid, text_value sql.NullString
                        var int_value sql.NullInt64
                        var bool_value sql.NullBool

			err := rows.Scan(&model.Id, &uid, &bool_value, &int_value, &text_value)
			if err != nil {
				log.Println(err)
				http.Error(w, http.StatusText(500), 500)
				return
			}

			log.Println(bool_value)

			model.Uid = uid.String
			model.Text = text_value.String
			model.Int = int(int_value.Int64)
			model.Bool = bool_value.Bool


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
