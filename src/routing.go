package main

import (
	"net/http"

	wishlist "github.com/Sam36502/WishlistLib-go"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {

	e.GET("/", pgMain)
	e.GET("/search", pgSearch)

	e.File("/favicon.ico", "data/img/favicon.ico")

}

type Data struct {
	CurrTime string
	System   string
}

func pgMain(c echo.Context) error {
	return c.Render(http.StatusOK, "main", nil)
}

func pgSearch(c echo.Context) error {
	ctx := wishlist.DefaultContext()
	users, _ := ctx.GetAllUsers()
	return c.Render(http.StatusOK, "search", users[0].Name)
}
