package main

import (
	"github.com/gaigepr/list-app/api"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"

	"net/http"
)

func main() {
	router := mux.NewRouter()
	InitHttpHandlers(router)

	api.InitializeDBConnection()
	defer api.DB.Close()

	http.ListenAndServe(":8080", context.ClearHandler(http.DefaultServeMux))
}
