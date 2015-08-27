package main

import (
	"database/sql"
	"net/http"
        _ "github.com/lib/pq"
        "log"
        "time"

)

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
