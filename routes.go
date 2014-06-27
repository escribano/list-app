package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func base(res http.ResponseWriter, req *http.Request) {}

func InitHttpHandlers(router *mux.Router) {
	routes := map[string]http.HandlerFunc{
		// Static file routes
		"/":         Index,
		"/login":    GetLogin,
		"/logout":   Index,
		"/register": GetNewAccount,

		// Routes related to users
		"/user/create": base,
		"/user/delete": base,
		"/user/update": base,

		// Routes related to tags
		"/tag/add": base,

		// Routes related to tasks
		"/task/create":          base,
		"/task/delete":          base,
		"/task/update/tags":     base,
		"/task/update/text":     base,
		"/task/update/deadline": base,
		"/task/get/all":         base,
		"/task/get/{id:[0-9]+}": base,
	}

	// Serve static directory
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	for route, handler := range routes {
		router.HandleFunc(route, handler)
	}

	http.Handle("/", router)
}
