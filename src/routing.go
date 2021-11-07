package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {

	e.GET("/", pgMain)
	e.GET("/search", pgSearch)
	e.GET("/register", pgRegister)
	e.POST("/register", registerUser)
	e.GET("/login", pgLogin)
	e.POST("/login", loginUser)
	e.GET("/user/:email", pgUserList)

	e.File("/favicon.ico", "data/img/favicon.ico")

}

type Data struct {
	CurrTime string
	System   string
}

func pgMain(c echo.Context) error {
	return c.Render(http.StatusOK, "main", nil)
}
