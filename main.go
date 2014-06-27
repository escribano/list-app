package main

import (
	_ "github.com/lib/pq"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	//"github.com/gorilla/rpc"
	//"github.com/gorilla/rpc/json"
	//"github.com/gorilla/schema"
	"github.com/gorilla/sessions"

	"database/sql"
	"fmt"
	"net/http"
)

var store = sessions.NewCookieStore([]byte("something-secret"))

func main() {
	connString := "postgres://" + DBUser + ":" + DBPass + "@localhost/" + DBName + "?sslmode=disable"
	db, err := sql.Open("postgres", connString)
	defer db.Close()
	if err != nil {
		fmt.Println("Connection error: ", err)
	}

	rows, err := db.Query("SELECT * FROM tags;")
	defer rows.Close()
	if err != nil {
		fmt.Println("Query error: ", err)
	}

	for rows.Next() {
		var pk int
		var tag string
		if err := rows.Scan(&pk, &tag); err != nil {
			fmt.Println("Row error: ", err)
		}
		fmt.Printf("IN row: %s %s\n", pk, tag)
	}

	router := mux.NewRouter()
	InitHttpHandlers(router)

	http.ListenAndServe(":8080", context.ClearHandler(http.DefaultServeMux))
}
