package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func InitHttpHandlers(router *mux.Router) {
	routes := map[string]http.HandlerFunc{
		"/":           IndexHandler,
		"/login":      LoginHandler,
		"/newaccount": NewAccountHandler,
	}

	// Serve static directory
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	for route, handler := range routes {
		if err := router.HandleFunc(route, handler); err != nil {
			fmt.Println(err)
		}
	}

	http.Handle("/", router)
}
