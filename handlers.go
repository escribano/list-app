package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func IndexHandler(res http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadFile("templates/index.html")
	fmt.Fprint(res, string(body))
}
