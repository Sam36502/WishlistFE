package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {

	e.GET("/", testPage)

}

func testPage(c echo.Context) error {
	return c.HTML(
		http.StatusOK,
		"<h1>It's working!</h2>",
	)
}
