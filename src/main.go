package main

import (
	//"os"

	"wishlist_fe/src/inf"

	"github.com/labstack/echo/v4"
)

func main() {

	// Initialise Echo Framework
	e := echo.New()
	inf.InitWishlistClient()
	inf.InitCookieStore()
	inf.LoadTemplates(e)
	InitRoutes(e)
	e.HTTPErrorHandler = inf.RedirectHTTPErrorHandler

	// For local debugging
	e.Logger.Fatal(e.Start(":5000"))
}
