package main

import (
	"net/http"
	"runtime"
	"time"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {

	e.GET("/", testPage)

}

type Data struct {
	CurrTime string
	System   string
}

func testPage(c echo.Context) error {
	return c.Render(
		http.StatusOK, "test",
		Data{
			CurrTime: time.Now().Format("Monday, 02-Jan-06 15:04:05 MST"),
			System:   runtime.GOOS,
		},
	)
}
