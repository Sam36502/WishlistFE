package handlers

import (
	"net/http"
	"wishlist_fe/src/inf"

	wishlistlib "github.com/Sam36502/WishlistLib-go"
	"github.com/labstack/echo/v4"
)

func PgItem(c echo.Context) error {

	email := c.Param("email")
	if email == "" {
		return echo.ErrNotFound
	}

	// Get Item information
	item, err := inf.GetItemFromPath(c)
	if err != nil {
		return nil
	}

	// Check if currently logged in as this user
	loggedIn := false
	liUser, err := inf.GetLoggedInUser(c)
	if err == nil {
		loggedIn = liUser.Email == email
	}

	// Check if user can reserve this item
	canReserve := !loggedIn && item.Status.StatusID == 1

	// Check if user can see the status of this item
	canSeeStatus := !loggedIn || item.Status.StatusID == 3

	return c.Render(http.StatusOK, "item", struct {
		Item         wishlistlib.Item
		LoggedIn     bool
		CanSeeStatus bool
		CanReserve   bool
		StatusColour string
		Email        string
	}{
		Item:         item,
		LoggedIn:     loggedIn,
		CanSeeStatus: canSeeStatus,
		CanReserve:   canReserve,
		StatusColour: inf.GetStatusColour(item.Status),
		Email:        email,
	})
}
