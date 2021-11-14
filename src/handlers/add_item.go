package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func PgNewItem(c echo.Context) error {
	return c.Render(http.StatusOK, "add_item", nil)
}

func NewItem(c echo.Context) error {
	return echo.ErrServiceUnavailable
}
