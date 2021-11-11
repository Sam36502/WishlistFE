package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func pgDeleteUser(c echo.Context) error {
	return c.Redirect(http.StatusMovedPermanently, "/")
}

func deleteUser(c echo.Context) error {
	return c.Redirect(http.StatusMovedPermanently, "/")
}
