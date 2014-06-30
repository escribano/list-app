package api

import (
	_ "github.com/lib/pq"

	"database/sql"
	"fmt"
)

var (
	DB      *sql.DB
	DBName  = ""
	DBUser  = ""
	DBPass  = ""
	SSLMode = ""
)

func InitializeDBConnection() {
	var err error
	connString := "postgres://" + DBUser + ":" + DBPass + "@localhost/" + DBName + "?sslmode=" + SSLMode
	DB, err = sql.Open("postgres", connString)
	if err != nil {
		fmt.Println("ERROR connection to DB: ")
		DB.Close()
		panic(err)
	}
}
