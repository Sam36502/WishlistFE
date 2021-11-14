package inf

import (
	"encoding/gob"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

const (
	COOKIE_TIMEOUT   = 60 * 60     // User data cookie expires after one hour
	COOKIE_FORM_DATA = "form-data" // Cookie name for sending error information to form pages
	COOKIE_USER_DATA = "user-data" // Cookie name for user login data
)

type CookieUser struct {
	Email    string
	Password string
}

var CookieStore *sessions.CookieStore

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
		MaxAge:   0,
		HttpOnly: false,
	}

	// Register Session Types
	gob.Register(FormError{})
	gob.Register(FormUser{})
	gob.Register(CookieUser{})
}
