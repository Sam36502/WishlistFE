/*
 *		USER LIST
 *
 *		All functions pertaining to displaying and
 *		editing the user's wishlist items
 *
 */

package main

import (
	"fmt"
	"net/http"

	wishlistlib "github.com/Sam36502/WishlistLib-go"
	"github.com/labstack/echo/v4"
)

func pgUserList(c echo.Context) error {

	// Get all the user's items
	email := c.Param("email")
	if email == "" {
		return c.Redirect(http.StatusPermanentRedirect, "/home")
	}

	wish := wishlistlib.Context{
		BaseUrl: WISHLIST_BASE_URL,
	}

	user, err := wish.GetUserByEmail(email)
	if err != nil {
		return c.Redirect(http.StatusPermanentRedirect, "/home")
	}
	items, err := wish.GetAllItems(user)
	if err != nil {
		fmt.Println("[ERROR] Failed to retrieve User's items:\n ", err)
		return c.Redirect(http.StatusPermanentRedirect, "/err/500.html")
	}

	// Check if currently logged in as this user
	loggedIn := false
	liUser, err := getLoggedInUser(c)
	if err == nil {
		loggedIn = user.Email == liUser.Email
	}

	return c.Render(http.StatusOK, "userlist", struct {
		User     wishlistlib.User
		Items    []wishlistlib.Item
		LoggedIn bool
	}{
		User:     user,
		Items:    items,
		LoggedIn: loggedIn,
	})
}
