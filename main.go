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
	conn, err := sql.Open("postgres", "user=list_app_user dbname=list_app  sslmode=verifyfull")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(conn)

	router := mux.NewRouter()
	InitHttpHandlers(router)

	http.ListenAndServe(":8080", context.ClearHandler(http.DefaultServeMux))
}
