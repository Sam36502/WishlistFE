package main

import (
	"net/http"
	"wishlist_fe/src/handlers"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {

	e.GET("/", PgMain)
	e.File("/favicon.ico", "data/img/favicon.ico")

	// User Routes
	e.GET("/search", handlers.PgSearch)

	e.GET("/register", handlers.PgRegister)
	e.POST("/register", handlers.RegisterUser)
	e.GET("/registersuccess", handlers.PgRegisterSuccess)

	e.GET("/login", handlers.PgLogin)
	e.POST("/login", handlers.LoginUser)

	e.GET("/user/:email", handlers.PgUserList)
	e.GET("/user/:email/delete", handlers.PgDeleteUser)
	e.POST("/user/:email/delete", handlers.DeleteUser)

	// Item Routes
	e.GET("/item/:item_id", handlers.PgItem)
	e.GET("/user/:email/newitem", handlers.PgNewItem)
	e.POST("/user/:email/newitem", handlers.NewItem)

	e.GET("/user/:email/delitem/:item_id", handlers.PgDelItem)
	e.POST("/user/:email/delitem/:item_id", handlers.DelItem)

}

type Data struct {
	CurrTime string
	System   string
}

func PgMain(c echo.Context) error {
	return c.Render(http.StatusOK, "main", nil)
}
