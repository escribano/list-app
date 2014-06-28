package main

import (
	_ "github.com/lib/pq"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	//"github.com/gorilla/rpc"
	//"github.com/gorilla/rpc/json"
	//"github.com/gorilla/schema"
	"github.com/gorilla/sessions"

	//"crypto/rand"
	"database/sql"
	"fmt"
	"net/http"
)

var (
	DB     *sql.DB
	DBName = "cool_db"
	DBUser = "cool_db_user"
	DBPass = "cool_passwrd!"

	Sessions *sessions.CookieStore = sessions.NewCookieStore([]byte("SUPER SECRET"))
)

func main() {
	connString := "postgres://" + DBUser + ":" + DBPass + "@localhost/" + DBName + "?sslmode=disable"
	DB, err := sql.Open("postgres", connString)
	if err != nil {
		fmt.Println("Connection error: ", err)
	}
	defer DB.Close()
	fmt.Println(DB)

	// rows, err := DB.Query("SELECT * FROM tags;")
	// defer rows.Close()
	// if err != nil {
	// 	fmt.Println("Query error: ", err)
	// }

	// for rows.Next() {
	// 	var pk int
	// 	var tag string
	// 	if err := rows.Scan(&pk, &tag); err != nil {
	// 		fmt.Println("Row error: ", err)
	// 	}
	// 	fmt.Printf("IN row: %s %s\n", pk, tag)
	// }

	router := mux.NewRouter()
	InitHttpHandlers(router)

	http.ListenAndServe(":8080", context.ClearHandler(http.DefaultServeMux))
}
