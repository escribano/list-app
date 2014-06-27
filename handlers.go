package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func IndexHandler(res http.ResponseWriter, req *http.Request) {
	// login := new(LoginForm)

	// err := req.ParseForm()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// err = decoder.Decode(login, req.PostForm)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// //call AuthPass
	// fmt.Println(login)

	body, _ := ioutil.ReadFile("templates/index.html")
	fmt.Fprint(res, string(body))
}

func NewAccountHandler(res http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadFile("templates/new-account.html")
	fmt.Fprint(res, string(body))
}

func LoginHandler(res http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadFile("templates/login.html")
	fmt.Fprint(res, string(body))
}
