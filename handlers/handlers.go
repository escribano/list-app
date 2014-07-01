package handlers

import (
	//"crypto/rand"
	"fmt"
	"net/http"
)

// Handler for getting all a user's tasks; apart of post login screen
func GetUserTasks(res http.ResponseWriter, req *http.Request) {
	//api.GetAllUserTasks
	session, err := Store.Get(req, "list-app")
	if err != nil {
		fmt.Println("ERROR gettting session: ", err)
	}
	fmt.Println("OLD SESSION: ", session)
}

// Handler for creating a new task if a user is authed
func PostNewTask(res http.ResponseWriter, req *http.Request) {
	//api.CreateTask
}

// Handler for updating a task's data
func PostUpdateTask(res http.ResponseWriter, req *http.Request) {
	//api.CreateNewTag and others
}

// Middleware stuff

// Use allows us to stack middleware to process the request
// Example taken from https://github.com/gorilla/mux/pull/36#issuecomment-25849172
func Use(handler http.HandlerFunc, mid ...func(http.Handler) http.HandlerFunc) http.HandlerFunc {
	for _, m := range mid {
		handler = m(handler)
	}
	return handler
}

// RequireLogin is a simple middleware which checks to see if the user is currently logged in.
// If not, the function returns a 302 redirect to the login page.
func RequireLogin(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if session, _ := Store.Get(r, "list-app"); session.Values["sessionId"] != nil {
			handler.ServeHTTP(w, r)
		} else {
			http.Redirect(w, r, "/login", 302)
		}
	}
}
