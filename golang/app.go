package main

import (
	"database/sql"
	"net/http"
        "encoding/json"
        _ "github.com/lib/pq"
        "log"
        "time"

)

type User struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

func psqlSelectHandler(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		rows, err := db.Query(`select id, email from accounts`)

		if err != nil {
			panic(err)
			http.Error(w, http.StatusText(500), 500)
		}

		defer rows.Close()

		users := make([]*User, 0)

		for rows.Next() {
			user := new(User)

			err := rows.Scan(&user.Id, &user.Email)
			if err != nil {
				panic(err)
			}

			users = append(users, user)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	})
}

func main() {

	db, err := sql.Open("postgres", "user=staging dbname=staging sslmode=disable port=5432 host=localhost password=")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(4)
	db.SetMaxOpenConns(5)
	defer db.Close()

	server := &http.Server{
		Addr:           ":8000",
		Handler:        nil,
		ReadTimeout:    100 * time.Second,
		WriteTimeout:   100 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	http.Handle("/psql/select", psqlSelectHandler(db))

	log.Println("Server listen on 8000")
	panic(server.ListenAndServe())

}
