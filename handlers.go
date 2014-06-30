package main

import (
	//"github.com/gorilla/rpc"
	//"github.com/gorilla/rpc/json"
	"github.com/gorilla/schema"

	"github.com/gaigepr/list-app/api"

	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	decoder = schema.NewDecoder()
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
	}
	if err := decoder.Decode(user, req.PostForm); err != nil {
		fmt.Println("ERROR decoding form: ", err)
	}

	if err := user.Validate(); err != nil {
		fmt.Println(err)
		fmt.Fprintf(res, err.Error())
		return //and redirect or something
	}

	userObj, err := api.GetUser(user.Email)
	if err != nil {
		fmt.Println("ERROR getting user: ", err)
	}

	if AuthPass(user.Password, userObj.Hash, userObj.Salt) {
		fmt.Fprint(res, "Success authenticating")
	} else {
		fmt.Fprint(res, "Failed to authenticate")
		return
	}

	// Make a new session with a random string as the name.
	// Save that string in the store or something so that none
	// of that data is client side.
	session, err := Sessions.Get(req, "session-name")
	fmt.Println(session.Name())
	if err != nil {
		fmt.Println(err)
	}

	if user.RememberMe {
		// Set session for 14 days
		//session.Options.MaxAge = 86400 * 14
		session.Options.MaxAge = 60
	} else {
		session.Options.MaxAge = 30
	}
	session.Save(req, res)
}
