package main

import (
	"github.com/gorilla/schema"

	"fmt"
	"io/ioutil"
	"net/http"
)

var decoder = schema.NewDecoder()

func Index(res http.ResponseWriter, req *http.Request) {
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

func GetNewAccount(res http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadFile("templates/new-account.html")
	fmt.Fprint(res, string(body))
}

func PostNewAccount(res http.ResponseWriter, req *http.Request) {
	newUser := new(CreateAccountForm)

	if err := req.ParseForm(); err != nil {
		fmt.Println("ERROR parsing form: ", err)
	}

	if err := decoder.Decode(newUser, req.PostForm); err != nil {
		// Could return an error here becuase it implies an emtpy field
		fmt.Println("ERROR decoding form: ", err)
	}

	var hash, salt string
	if newUser.Password == newUser.Password2 {
		hash, salt = NewPass(newUser.Password)
	}

	queryString := fmt.Sprintf(`INSERT INTO users (first_name, last_name, email, hash, salt) VALUES ("%s", "%s", "%s", E"%s", E"%s");`,
		newUser.FirstName, newUser.LastName, newUser.Email, hash, salt)
	fmt.Println(queryString)
	results, err := DB.Exec(queryString)
	if err != nil {
		fmt.Println("ERROR inserting new user: ", err)
	}
	fmt.Println(results)

	fmt.Fprint(res, newUser)
}

func GetLogin(res http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadFile("templates/login.html")
	fmt.Fprint(res, string(body))
}

func PostLogin(res http.ResponseWriter, req *http.Request) {

}

func GetAllTasks(res http.ResponseWriter, req *http.Request) {

}
