package main

import (
	"os"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

const (
	WISHLIST_BASE_URL = "https://www.pearcenet.ch:2512"
	COOKIE_TIMEOUT    = 60 * 60     // User data cookie expires after one hour
	COOKIE_FORM_DATA  = "form-data" // Cookie name for sending error information to form pages
	COOKIE_USER_DATA  = "user-data" // Cookie name for user login data
)

var CookieStore *sessions.CookieStore

func main() {

	// Initialise Echo Framework
	e := echo.New()
	InitCookieStore()
	InitRoutes(e)
	LoadTemplates(e)

	// For local debugging
	//e.Logger.Fatal(e.Start(":5000"))

	e.Logger.Fatal(e.StartTLS(":"+os.Getenv("WISHLIST_FE_PORT"), os.Getenv("WISHLIST_SSL_CERT"), os.Getenv("WISHLIST_SSL_KEY")))
}
