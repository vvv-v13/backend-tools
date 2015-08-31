package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"time"
)

func main() {

        // PostgreSQL Connection pool
	db, err := sql.Open("postgres", "user=test dbname=test sslmode=disable port=5432 host=localhost password=")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(4)
	db.SetMaxOpenConns(5)
	defer db.Close()

        // Http server
	server := &http.Server{
		Addr:           ":8000",
		Handler:        nil,
		ReadTimeout:    100 * time.Second,
		WriteTimeout:   100 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

        // Routing
	http.Handle("/psql/select", psqlSelectHandler(db))
	http.Handle("/psql/insert", psqlInsertHandler(db))

	log.Println("Server listen on 8000")
	panic(server.ListenAndServe())

}
