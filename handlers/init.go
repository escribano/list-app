package handlers

import (
	"github.com/gorilla/schema"
	"github.com/gorilla/sessions"

	"crypto/rand"
)

var (
	hashKey  = make([]byte, 64)
	blockKey = make([]byte, 32)
	Decoder  *schema.Decoder
	Store    *sessions.CookieStore
)

func init() {
	rand.Read(hashKey)
	rand.Read(blockKey)
	Store = sessions.NewCookieStore(hashKey, blockKey)
	Decoder = schema.NewDecoder()
}
