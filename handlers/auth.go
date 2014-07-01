package handlers

import (
	"code.google.com/p/go.crypto/pbkdf2"
	"github.com/gaigepr/list-app/api"

	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"net/http"
)

func NewPass(password string) (string, string) {
	// make user salt and hash of pass and put into database
	// for salt may want to have seperate database that stores all salts ever made
	// to make sure none are reused.
	//salt := []byte("fixed_test_salts")
	//fmt.Println(len(salt))
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		fmt.Println(err)
	}
	return base64.StdEncoding.EncodeToString(pbkdf2.Key([]byte(password), salt, 4096, 64, sha512.New)), base64.StdEncoding.EncodeToString(salt)
}

func AuthPass(password, hash, salt string) bool {
	userSalt, err := base64.StdEncoding.DecodeString(salt)
	if err != nil {
		fmt.Println("error de-base64-ing salt: ", err)
	}
	userHash, err := base64.StdEncoding.DecodeString(hash)
	if err != nil {
		fmt.Println("error de-base64-ing hash: ", err)
	}

	//hash potential input password
	passHash := pbkdf2.Key([]byte(password), userSalt, 4096, 64, sha512.New)

	// compare passHash to the hash in the database,
	// if match login is correct(true), if not login failed(false).

	// my (probably wrong) attempt at constant time comparison to defeat timing attacks
	// if the hashes are different sized iterate through the bigger one anyway
	// if the hashes are same size iterate through, if any dont match up set a bool to false
	// at end of matching check bool value.
	match := true
	for i := 0; i < len(userHash); i++ {
		if userHash[i] != passHash[i] {
			match = false
		}
	}
	return match
}

// Handler for making a new user
func PostNewAccount(res http.ResponseWriter, req *http.Request) {
	newUser := new(CreateAccountForm)

	if err := req.ParseForm(); err != nil {
		fmt.Println("ERROR parsing form: ", err)
	}
	if err := Decoder.Decode(newUser, req.PostForm); err != nil {
		fmt.Println("ERROR decoding form: ", err)
	}

	if err := newUser.Validate(); err != nil {
		fmt.Println(err)
		fmt.Fprintf(res, err.Error())
		return
	}

	hash, salt := NewPass(newUser.Password)
	if err := api.CreateNewUser(newUser.Email, newUser.FirstName, newUser.LastName, hash, salt); err != nil {
		fmt.Println("Error creating new user: ", err) // Make a log function bruh
		fmt.Fprint(res, "There was an error creating your account.")
		return
	}

	fmt.Println("Made a new user: ", newUser, hash, salt)
	fmt.Fprint(res, "Booyah, success!")
}

// Handler for authenticating a user
func PostLogin(res http.ResponseWriter, req *http.Request) {
	user := new(LoginForm)

	if err := req.ParseForm(); err != nil {
		fmt.Println("ERROR parsing form: ", err)
		fmt.Fprint(res, "<h1>500: Internal Server Error</h1>")
		return
	}
	if err := Decoder.Decode(user, req.PostForm); err != nil {
		fmt.Println("ERROR decoding form: ", err)
		fmt.Fprint(res, "<h1>500: Internal Server Error</h1>")
		return
	}

	if err := user.Validate(); err != nil {
		fmt.Println(err)
		fmt.Fprintf(res, err.Error())
		return //and redirect or something
	}

	userObj, err := api.GetUser(user.Email)
	if err != nil {
		fmt.Println("ERROR getting user: ", err)
		fmt.Fprint(res, "<h1>500: Internal Server Error</h1>")
		return
	}

	if !AuthPass(user.Password, userObj.Hash, userObj.Salt) {
		fmt.Fprint(res, "Failed to authenticate")
		return
	}

	// Make a new session with a random string as the name.
	// Save that string in the Store or something so that none
	// of that data is client side.
	session, err := Store.Get(req, "list-app")
	if err != nil {
		fmt.Println("ERROR gettting session: ", err)
	}
	// generate and insert a unique and long session id
	session.Values["sessionId"] = "FUCKING GEEZUZ"

	if user.RememberMe {
		session.Options.MaxAge = 120
	} else {
		session.Options.MaxAge = 30
	}

	if err := session.Save(req, res); err != nil {
		fmt.Println("ERROR saving session: ", err)
	}
	fmt.Println("SAVED: ", session)
	http.Redirect(res, req, "/task/get/all", 302)
}
