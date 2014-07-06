package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
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

// list-view.html page
func DevList(res http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadFile("templates/list-view-v2.html")
	fmt.Fprint(res, string(body))
}

// list-view.html page
func Dev(res http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadFile("templates/dev.html")
	fmt.Fprint(res, string(body))
}
