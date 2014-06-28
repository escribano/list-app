package main

import (
	//"github.com/gorilla/rpc"
	//"github.com/gorilla/rpc/json"
	"github.com/gorilla/schema"

	"fmt"
	"io/ioutil"
	"net/http"
)

var decoder = schema.NewDecoder()

// Get index.html form
func Index(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Cookie("session-name"))
	body, _ := ioutil.ReadFile("templates/index.html")
	fmt.Fprint(res, string(body))
}

// Get register.html form
func GetNewAccount(res http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadFile("templates/register.html")
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

	var hash, salt string
	if newUser.Password == newUser.Password2 {
		hash, salt = NewPass(newUser.Password)
	}

	// API call here? api.InsertNewUser()

	fmt.Println(newUser, hash, salt)
	fmt.Fprint(res, newUser)
}

// Get login.html form
func GetLogin(res http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadFile("templates/login.html")
	fmt.Fprint(res, string(body))
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

	fmt.Println("AUTHED USER: ", user)
	// get user salt and hash to validate
	// if ok, redirect, otherwise rety login

	// Make a new session with a random string as the name.
	// Save that string in the store or something so that none
	// of that data is client side.
	session, err := Sessions.Get(req, "session-name")
	fmt.Println(session.Name())
	if err != nil {
		fmt.Println(err)
	}

	session.Values["authenticated"] = true
	//session.Values["id"] = API CALL
	if user.RememberMe {
		// Set session for 14 days
		//session.Options.MaxAge = 86400 * 14
		session.Options.MaxAge = 60
	} else {
		session.Options.MaxAge = 30
	}
	session.Save(req, res)

}
