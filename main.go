package main

import (
	"database/sql"
	_ "github.com/lib/pq"

	"fmt"
)

func main() {
	db, err := sql.Open("postgres", "user=list_app dbname=list_app  sslmode=verifyfull")
	if err != nil {
		fmt.Println(err)
	}

}
