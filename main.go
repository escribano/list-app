package main

import (
	"database/sql"
	_ "github.com/lib/pq"

	"github.com/gorilla/mux"

	"fmt"
	"net/http"
)

func index(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "<h1>HI</h1>")
}

func main() {
	_, err := sql.Open("postgres", "user=list_app dbname=list_app  sslmode=verifyfull")
	if err != nil {
		fmt.Println(err)
	}

	router := mux.NewRouter()

	// Serve static directory
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	router.HandleFunc("/", index)
	router.HandleFunc("/{name}", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "<h1>Hello there, Mc-%s</h1>", req.URL.Path[1:])
	})

	http.Handle("/", router)
	http.ListenAndServe(":8088", nil)
}
