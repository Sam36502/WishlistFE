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
	e.GET("/logout", handlers.Logout)

	e.GET("/user/:email", handlers.PgUserList)
	e.GET("/user/:email/delete", handlers.PgDeleteUser)
	e.POST("/user/:email/delete", handlers.DeleteUser)

	// Item Routes
	e.GET("/user/:email/item/:item_id", handlers.PgItem)
	e.GET("/user/:email/newitem", handlers.PgNewItem)
	e.POST("/user/:email/newitem", handlers.NewItem)

	e.GET("/user/:email/:item_id/delete", handlers.PgDelItem)
	e.POST("/user/:email/:item_id/delete", handlers.DelItem)

	e.GET("/user/:email/item/:item_id/reserve", handlers.PgReserveItem)
	e.POST("/user/:email/item/:item_id/reserve", handlers.ReserveItem)
	e.GET("/user/:email/item/:item_id/unreserve", handlers.PgUnreserveItem)
	e.POST("/user/:email/item/:item_id/unreserve", handlers.UnreserveItem)
	e.GET("/user/:email/item/:item_id/receive", handlers.PgReceiveItem)
	e.POST("/user/:email/item/:item_id/receive", handlers.ReceiveItem)
	e.GET("/user/:email/item/:item_id/unreceive", handlers.PgUnReceiveItem)
	e.POST("/user/:email/item/:item_id/unreceive", handlers.UnReceiveItem)

}

func PgMain(c echo.Context) error {
	return c.Render(http.StatusOK, "main", nil)
}
