/*
 *		SEARCH
 *
 *		All functions pertaining to searching for users
 *
 */

package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"wishlist_fe/src/inf"

	wishlistlib "github.com/Sam36502/WishlistLib-go"
	"github.com/labstack/echo/v4"
)

type searchResultData struct {
	Search       string
	ResultString string
	Results      []wishlistlib.User
}

func PgSearch(c echo.Context) error {

	// Get Search Query
	if !c.QueryParams().Has("s") {
		return c.Render(http.StatusOK, "search", searchResultData{
			Search:       "",
			ResultString: "Enter a name above to find friends.",
			Results:      nil,
		})
	}
	search := c.QueryParam("s")
	search = strings.ToLower(search)
	search = strings.TrimSpace(search)

	// Search Users
	wish := inf.GetWishlistClient(wishlistlib.Token{})
	users, err := wish.SearchUsers(search)
	if err != nil {
		fmt.Println("[ERROR] Failed to search users:\n ", err)
		return echo.ErrInternalServerError
	}

	resString := strconv.Itoa(len(users)) + " Results:"
	switch len(users) {
	case 0:
		resString = "No Results"
	case 1:
		resString = "1 Result:"
	}

	return c.Render(http.StatusOK, "search", searchResultData{
		Search:       search,
		ResultString: resString,
		Results:      users,
	})
}
