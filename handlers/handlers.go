package handlers

import (
	"github.com/gaigepr/list-app/api"

	//"crypto/rand"
	"encoding/json"
	"fmt"
	"net/http"
)

type JsonResponse struct {
	Name string
	Data interface{}
}

// Handler for getting all a user's tasks; apart of post login screen
func GetUserTasks(res http.ResponseWriter, req *http.Request) {
	session, err := Store.Get(req, "list-app")
	if err != nil {
		fmt.Println("ERROR gettting session: ", err)
		fmt.Fprint(res, "500: Internal Server Error.")
		return
	}

	tasks, err := api.GetAllUserTasks(session.Values["userId"].(int))
	fmt.Println(tasks)
	if err != nil {
		fmt.Println(err)
	}

	b, err := json.Marshal(JsonResponse{"GetUserTasks", tasks})
	if err != nil {
		fmt.Println("ERROR marshaling", err)
	}

	fmt.Fprint(res, (string(b)))
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
