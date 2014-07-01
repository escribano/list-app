package handlers

import (
	//"github.com/gorilla/rpc"
	//"github.com/gorilla/rpc/js
	"github.com/gorilla/schema"
	"github.com/gorilla/sessions"

	"github.com/gaigepr/list-app/api"

	//"crypto/rand"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	hashKey  = []byte("a-very-secret-thing-omg")
	blockKey = []byte("we-must-encrypt-all-of-the-thing")
	store    = sessions.NewCookieStore([]byte(" buttmang"), []byte("12345678901234567890123456789012"))
	decoder  = schema.NewDecoder()
)

// Get index.html form
func Index(res http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadFile("templates/index.html")
	fmt.Fprint(res, string(body))
}

// Get register.html form
func GetNewAccount(res http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadFile("templates/register.html")
	fmt.Fprint(res, string(body))
}

// Get login.html form
func GetLogin(res http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadFile("templates/login.html")
	fmt.Fprint(res, string(body))
}

// list-view.html page
func GetListView(res http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadFile("templates/list-view.html")
	fmt.Fprint(res, string(body))
}

// Handler for making a new user
func PostNewAccount(res http.ResponseWriter, req *http.Request) {
	newUser := new(CreateAccountForm)

	if err := req.ParseForm(); err != nil {
		fmt.Println("ERROR parsing form: ", err)
	}
	if err := decoder.Decode(newUser, req.PostForm); err != nil {
		fmt.Println("ERROR decoding form: ", err)
	}

	if err := newUser.Validate(); err != nil {
		fmt.Println(err)
		fmt.Fprintf(res, err.Error())
		return
	}

	hash, salt := NewPass(newUser.Password)
	if err := api.CreateNewUser(newUser.Email, newUser.FirstName, newUser.LastName, hash, salt); err != nil {
		fmt.Println("Error creating new user: ", err) // Make a log function bruh
		fmt.Fprint(res, "There was an error creating your account.")
		return
	}

	fmt.Println("Made a new user: ", newUser, hash, salt)
	fmt.Fprint(res, "Booyah, success!")
}

// Handler for authenticating a user
func PostLogin(res http.ResponseWriter, req *http.Request) {
	user := new(LoginForm)

	if err := req.ParseForm(); err != nil {
		fmt.Println("ERROR parsing form: ", err)
		fmt.Fprint(res, "<h1>500: Internal Server Error</h1>")
		return
	}
	if err := decoder.Decode(user, req.PostForm); err != nil {
		fmt.Println("ERROR decoding form: ", err)
		fmt.Fprint(res, "<h1>500: Internal Server Error</h1>")
		return
	}

	if err := user.Validate(); err != nil {
		fmt.Println(err)
		fmt.Fprintf(res, err.Error())
		return //and redirect or something
	}

	userObj, err := api.GetUser(user.Email)
	if err != nil {
		fmt.Println("ERROR getting user: ", err)
		fmt.Fprint(res, "<h1>500: Internal Server Error</h1>")
		return
	}

	if !AuthPass(user.Password, userObj.Hash, userObj.Salt) {
		fmt.Fprint(res, "Failed to authenticate")
		return
	}

	// Make a new session with a random string as the name.
	// Save that string in the store or something so that none
	// of that data is client side.
	session, err := store.Get(req, "list-app")
	if err != nil {
		fmt.Println("ERROR gettting session: ", err)
	}
	// generate and insert a unique and long session id
	session.Values["sessionId"] = "FUCKING GEEZUZ"

	if user.RememberMe {
		session.Options.MaxAge = 120
	} else {
		session.Options.MaxAge = 30
	}

	if err := session.Save(req, res); err != nil {
		fmt.Println("ERROR saving session: ", err)
	}
	fmt.Println("SAVED: ", session)
	http.Redirect(res, req, "/task/get/all", 302)
}

// Handler for getting all a user's tasks; apart of post login screen
func GetUserTasks(res http.ResponseWriter, req *http.Request) {
	//api.GetAllUserTasks
	session, err := store.Get(req, "list-app")
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
		if session, _ := store.Get(r, "list-app"); session.Values["sessionId"] != nil {
			handler.ServeHTTP(w, r)
		} else {
			http.Redirect(w, r, "/login", 302)
		}
	}
}
