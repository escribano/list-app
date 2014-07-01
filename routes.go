package main

import (
	. "github.com/gaigepr/list-app/handlers"
	"github.com/gorilla/mux"

	"net/http"
)

func base(res http.ResponseWriter, req *http.Request) {}

func InitHttpHandlers(router *mux.Router) {
	routes := map[string]http.HandlerFunc{
		// Static file routes
		"/":          Index,
		"/login":     GetLogin,
		"/register":  GetNewAccount,
		"/list-view": GetListView, // temp route

		// Routes related to users
		"/user/create":       PostNewAccount,
		"/user/delete":       base,
		"/user/update":       base,
		"/user/authenticate": PostLogin,
		"/user/logout":       base,

		// Routes related to tags
		"/tag/add": base,

		// Routes related to tasks
		"/task/create":          base,
		"/task/delete":          base,
		"/task/update":          base,
		"/task/get/all":         Use(GetUserTasks, RequireLogin),
		"/task/get/{id:[0-9]+}": base,
	}

	// Serve static directory
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	for route, handler := range routes {
		router.HandleFunc(route, handler)
	}

	http.Handle("/", router)
}
