package main

import (
	"net/http"

	"github.com/gorilla/mux"
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
		"/task/update/tags":     base,                            // Perhaps make one update routes and pass the relevant info in json
		"/task/update/text":     base,                            // Update teh info for a task
		"/task/update/deadline": base,                            // Set a dealine for a task
		"/task/get/all":         Use(GetUserTasks, RequireLogin), // Default route to be used on login
		"/task/get/{id:[0-9]+}": base,                            // This may not be a useful route.
	}

	// Serve static directory
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	for route, handler := range routes {
		router.HandleFunc(route, handler)
	}

	http.Handle("/", router)
}
