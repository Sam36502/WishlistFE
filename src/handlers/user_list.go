/*
 *		USER LIST
 *
 *		All functions pertaining to displaying and
 *		editing the user's wishlist items
 *
 */

package handlers

import (
	"fmt"
	"net/http"
	"wishlist_fe/src/inf"

	wishlistlib "github.com/Sam36502/WishlistLib-go"
	"github.com/labstack/echo/v4"
)

func PgUserList(c echo.Context) error {

	// Get all the user's items
	email := c.Param("email")
	if email == "" {
		return echo.ErrNotFound
	}

	wish := wishlistlib.DefaultWishClient(inf.WISHLIST_BASE_URL)

	user, err := wish.GetUserByEmail(email)
	if err != nil {
		return c.Redirect(http.StatusPermanentRedirect, "/home")
	}
	items, err := wish.GetAllItemsOfUser(user)
	if err != nil {
		fmt.Println("[ERROR] Failed to retrieve User's items:\n ", err)
		return c.Redirect(http.StatusPermanentRedirect, "/err/500.html")
	}

	// Check if currently logged in as this user
	loggedIn := false
	liUser, _, err := inf.GetLoggedInUser(c)
	fmt.Println("[DEBUG] Err from fetching logged-in user:", err)
	if err == nil {
		loggedIn = user.Email == liUser.Email
	}

	return c.Render(http.StatusOK, "user_list", struct {
		User     wishlistlib.User
		Items    []wishlistlib.Item
		LoggedIn bool
	}{
		User:     user,
		Items:    items,
		LoggedIn: loggedIn,
	})
}
