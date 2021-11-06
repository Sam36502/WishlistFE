package main

import (
	"encoding/gob"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

func InitCookieStore() {
	// Initialise Cookie Store
	authKeyOne := securecookie.GenerateRandomKey(64)
	encryptionKeyOne := securecookie.GenerateRandomKey(32)

	CookieStore = sessions.NewCookieStore(
		authKeyOne,
		encryptionKeyOne,
	)

	CookieStore.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   60 * 15,
		HttpOnly: false,
	}

	// Register Session Types
	gob.Register(FormError{})
	gob.Register(FormUser{})
}
