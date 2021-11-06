package main

import (
	"os"

	wishlistlib "github.com/Sam36502/WishlistLib-go"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

var CookieStore *sessions.CookieStore
var Wishlist *wishlistlib.Context

func main() {

	// Initialise Echo Framework
	e := echo.New()
	Wishlist = wishlistlib.DefaultContext()
	Wishlist.BaseUrl = "https://www.pearcenet.ch:2512"
	InitCookieStore()
	InitRoutes(e)
	LoadTemplates(e)

	// For local debugging
	//e.Logger.Fatal(e.Start(":5000"))

	e.Logger.Fatal(e.StartTLS(":"+os.Getenv("WISHLIST_FE_PORT"), os.Getenv("WISHLIST_SSL_CERT"), os.Getenv("WISHLIST_SSL_KEY")))
}
