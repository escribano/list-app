package main

import (
	//"github.com/gorilla/context"
	//"github.com/gorilla/mux"
	//"github.com/gorilla/rpc"
	//"github.com/gorilla/rpc/json"
	"github.com/gorilla/schema"
	//"github.com/gorilla/sessions"
	"crypto/sha512"
	"crypto/rand"
	"code.google.com/p/go.crypto/pbkdf2"
	"fmt"
	"net/http"
)

type CreateAccountForm struct {
	FirstName string
	LastName  string
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
	//call AuthPass
	fmt.Println(login)
}

func NewAccountHandler(res http.ResponseWriter, req *http.Request) {
	newUser := new(CreateAccountForm)

	err := req.ParseForm()
	if err != nil {
		fmt.Println(err)
	}

	err = decoder.Decode(newUser, req.PostForm)
	if err != nil {
		fmt.Println(err)
	}
	//call NewPass

	fmt.Println(newUser)
}

func NewPass(password string) ([]byte, []byte) {
	// make user salt and hash of pass and put into database
	// for salt may want to have seperate database that stores all salts ever made
	// to make sure none are reused.
	salt := []byte("fixed_test_salts")
	fmt.Println(len(salt))
	//salt := make([]byte, 16)
	//_, err := rand.Read(salt)
	//if err != nil {
	//fmt.Println(err)
	//}
	return pbkdf2.Key([]byte(password), salt, 4096, 64, sha512.New), salt
}


func AuthPass(password, user string) bool {
	// pull user's salt and password from database
	userSalt := []byte("fixed_test_salts")
	//hash potential input password
	passHash := pbkdf2.Key([]byte(login.Password), userSalt, 4096, 64, sha512.New)
	// compare passHash to the hash in the database, if match login is correct(true), if not login failed(false).
	return true
}
