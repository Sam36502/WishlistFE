package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func PgItem(c echo.Context) error {

	item_id := c.Param("item_id")
	if item_id == "" {
		return c.NoContent(http.StatusNotFound)
	}

	return c.Render(http.StatusOK, "item", nil)
}
