package main

import (
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	InitRoutes(e)
	LoadTemplates(e)

	// For local debugging
	//e.Logger.Fatal(e.Start(":5000"))

	e.Logger.Fatal(e.StartTLS(":"+os.Getenv("WISHLIST_FE_PORT"), os.Getenv("WISHLIST_SSL_CERT"), os.Getenv("WISHLIST_SSL_KEY")))
}
