package inf

import (
	"encoding/gob"

	wishlistlib "github.com/Sam36502/WishlistLib-go"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

type TokenJWTClaims struct {
	jwt.StandardClaims
	Email string `json:"email"`
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
	gob.Register(UserFormError{})
	gob.Register(FormUser{})
	gob.Register(ItemFormError{})
	gob.Register(FormItem{})
	gob.Register(wishlistlib.Token{})
}
