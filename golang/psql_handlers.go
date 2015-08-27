package main

import (
        "database/sql"
        "net/http"
        "encoding/json"
        _ "github.com/lib/pq"

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
