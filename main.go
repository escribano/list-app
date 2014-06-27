package main

import (
	"database/sql"
	_ "github.com/lib/pq"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	//"github.com/gorilla/rpc"
	//"github.com/gorilla/rpc/json"
	//"github.com/gorilla/schema"
	"github.com/gorilla/sessions"

	"fmt"
	"net/http"
)

var store = sessions.NewCookieStore([]byte("something-secret"))

func main() {
	_, err := sql.Open("postgres", "user=list_app dbname=list_app  sslmode=verifyfull")
	if err != nil {
		fmt.Println(err)
	}

	router := mux.NewRouter()

	// Serve static directory
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	router.HandleFunc("/login", LoginHandler)
	router.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		session, _ := store.Get(req, "session-name")
		session.Values["age"] = "20"
		err := session.Save(req, res)
		if err != nil {
			fmt.Println(err)
		}

		session.AddFlash("<h1>hello there flashee</h1>")
		if flashes := session.Flashes(); len(flashes) > 0 {
			fmt.Fprintf(res, "%s", flashes[0])
			session.Values["name"] = "dude"
		}

		session.Save(req, res)

		fmt.Fprintf(res, "\n<h1>Hi, %s. You're %s</h1>", session.Values["name"], session.Values["age"])
	})
	// THis should be protected
	router.HandleFunc("/derp", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, `<html><body><form method="post" action="/login">
                                 <p><input type="text" name="Username" value="" placeholder="Username or Email"></p>
                                 <p><input type="password" name="Password" value="" placeholder="Password"></p>
                                 <p class="remember_me">
                                 <label>
                                 <input type="checkbox" name="RememberMe"remember_me">
                                 Remember me on this computer
                                 </label>
                                 </p>
                                 <p class="submit"><input type="submit" name="" value="Login"></p>
                                 </form></body></html>`)
	})

	http.Handle("/", router)
	http.ListenAndServe(":8080", context.ClearHandler(http.DefaultServeMux))
}
