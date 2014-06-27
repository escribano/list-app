package main

import (
	//"github.com/gorilla/context"
	//"github.com/gorilla/mux"
	//"github.com/gorilla/rpc"
	//"github.com/gorilla/rpc/json"
	"github.com/gorilla/schema"
	//"github.com/gorilla/sessions"

	"fmt"
	"net/http"
)

type CreateAccountForm struct {
	FirstName string
	LastName  String
	Username  string
	Email     string
	Password  string
}

type LoginForm struct {
	Username   string
	Password   string
	RememberMe bool
}

var decoder = schema.NewDecoder()

func LoginHandler(res http.ResponseWriter, req *http.Request) {
	login := new(LoginForm)

	err := req.ParseForm()
	if err != nil {
		fmt.Println(err)
	}

	err = decoder.Decode(login, req.PostForm)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(login)
}

func NewAccountHandler(res http.RequestWriter, req *http.Request) {
	newUser := new(CreateAccountForm)

	err := req.ParseForm()
	if err != nil {
		fmt.Println(err)
	}

	err = decoder.Decode(newUser, req.PostForm)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(newUser)
}
