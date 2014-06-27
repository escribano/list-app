package main

import (
	//"github.com/gorilla/context"
	//"github.com/gorilla/mux"
	//"github.com/gorilla/rpc"
	//"github.com/gorilla/rpc/json"
	//"github.com/gorilla/schema"
	//"github.com/gorilla/sessions"
	"code.google.com/p/go.crypto/pbkdf2"
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	//"net/http"
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

func AuthPass(user, password string) bool {
	// pull user's salt and password from database
	rows, err := DB.Query("SELECT salt, hash FROM users WHERE user=" + user + ";")
	defer rows.Close()
	if err != nil {
		fmt.Println("Query error: ", err)
	}
	if !rows.Next() {
		fmt.Println("Error, no valid results returned")
	}
	var dbSalt string
	var dbHash string
	err = rows.Scan(&dbSalt, &dbHash)
	if err != nil {
		fmt.Println("Scanning error: ", err)
	}

	userSalt, err := base64.StdEncoding.DecodeString(dbSalt)
	if err != nil {
		fmt.Println("error de-base64-ing salt: ", err)
	}
	userHash, err := base64.StdEncoding.DecodeString(dbHash)
	if err != nil {
		fmt.Println("error de-base64-ing hash: ", err)
	}

	//hash potential input password
	passHash := pbkdf2.Key([]byte(user), userSalt, 4096, 64, sha512.New)

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
