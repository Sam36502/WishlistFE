package handlers

import (
	"net/http"
	"strconv"
	"wishlist_fe/src/inf"

	wishlistlib "github.com/Sam36502/WishlistLib-go"
	"github.com/labstack/echo/v4"
)

func PgDelItem(c echo.Context) error {

	email := c.Param("email")
	if email == "" {
		return echo.ErrNotFound
	}

	// Get item info
	var item_id uint64
	var err error
	if item_id_str := c.Param("item_id"); item_id_str == "" {
		return echo.ErrNotFound
	} else {
		item_id, err = strconv.ParseUint(item_id_str, 10, 64)
		if err != nil {
			return echo.ErrNotFound
		}
	}
	wish := wishlistlib.Context{
		BaseUrl: inf.WISHLIST_BASE_URL,
	}
	item, err := wish.GetItemByID(item_id)
	if err != nil {
		return echo.ErrNotFound
	}

	return c.Render(http.StatusOK, "del_item", struct {
		YesURL string
		NoURL  string
		Item   wishlistlib.Item
	}{
		YesURL: "/user/" + email + "/delitem/" + strconv.FormatUint(item_id, 10),
		NoURL:  "/user/" + email,
		Item:   item,
	})
}

func DelItem(c echo.Context) error {

	email := c.Param("email")
	if email == "" {
		return echo.ErrNotFound
	}

	// Get item info
	var item_id uint64
	var err error
	if item_id_str := c.Param("item_id"); item_id_str == "" {
		return echo.ErrNotFound
	} else {
		item_id, err = strconv.ParseUint(item_id_str, 10, 64)
		if err != nil {
			return echo.ErrNotFound
		}
	}
	wish := wishlistlib.Context{
		BaseUrl: inf.WISHLIST_BASE_URL,
	}
	item, err := wish.GetItemByID(item_id)
	if err != nil {
		return echo.ErrNotFound
	}

	// Check if currently logged in as this user
	liUser, err := inf.GetLoggedInUser(c)
	if err == nil {
		if email != liUser.Email {
			return echo.ErrForbidden
		}
	} else {
		return echo.ErrUnauthorized
	}

	// Delete item
	wish.SetAuthenticatedUser(liUser)
	err = wish.DeleteItem(item)
	if err != nil {
		return c.Render(http.StatusOK, "status", inf.StatusPageData{
			Colour:          "red",
			MainMessage:     "Failed to delete item. Please try again later",
			NextPageURL:     "/user/" + email,
			NextPageMessage: "Back to user page",
		})
	}

	return c.Render(http.StatusOK, "status", inf.StatusPageData{
		Colour:          "green",
		MainMessage:     "Item successfully deleted!",
		NextPageURL:     "/user/" + email,
		NextPageMessage: "Back to user page",
	})
}
