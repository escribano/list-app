package main

import (
	_ "github.com/lib/pq"

	"github.com/gaigepr/list-app/api"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	//"github.com/gorilla/rpc"
	//"github.com/gorilla/rpc/json"
	//"github.com/gorilla/schema"
	"github.com/gorilla/sessions"

	//"crypto/rand"
	//"database/sql"
	"fmt"
	"net/http"
)

var (
	Sessions *sessions.CookieStore = sessions.NewCookieStore([]byte("SUPER SECRET"))
)

func main() {
	fmt.Println()
	router := mux.NewRouter()
	InitHttpHandlers(router)

	api.InitializeDBConnection()
	defer api.DB.Close()
	api.GetUser("gaige@chatsubo.net")

	http.ListenAndServe(":8080", context.ClearHandler(http.DefaultServeMux))
}
