package main

import (
	"os"
	"wishlist_fe/src/inf"

	"github.com/labstack/echo/v4"
)

func main() {

	// Initialise Echo Framework
	e := echo.New()
	inf.InitCookieStore()
	inf.LoadTemplates(e)
	InitRoutes(e)
	e.HTTPErrorHandler = inf.RedirectHTTPErrorHandler

	// For local debugging
	//e.Logger.Fatal(e.Start(":5000"))

	e.Logger.Fatal(e.StartTLS(":"+os.Getenv("WISHLIST_FE_PORT"), os.Getenv("WISHLIST_SSL_CERT"), os.Getenv("WISHLIST_SSL_KEY")))
}
