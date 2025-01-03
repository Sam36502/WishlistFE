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

	wish := inf.GetWishlistClient(wishlistlib.Token{})

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
	thisIsMe := false
	liUser, _, err := inf.GetLoggedInUser(c)
	if err == nil {
		thisIsMe = user.Email == liUser.Email
	}

	return c.Render(http.StatusOK, "user_list", struct {
		User            wishlistlib.User
		Items           []wishlistlib.Item
		GetStatusColour func(wishlistlib.Status) string
		ThisIsMe        bool
	}{
		User:            user,
		Items:           items,
		GetStatusColour: inf.GetStatusColour,
		ThisIsMe:        thisIsMe,
	})
}
