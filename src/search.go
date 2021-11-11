/*
 *		SEARCH
 *
 *		All functions pertaining to searching for users
 *
 */

package main

import (
	"net/http"
	"strconv"
	"strings"

	wishlist "github.com/Sam36502/WishlistLib-go"
	"github.com/labstack/echo/v4"
)

type searchResultData struct {
	Search       string
	ResultString string
	Results      []wishlist.User
}

// Request Handler
func pgSearch(c echo.Context) error {

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
	ctx := wishlist.DefaultContext()
	users, err := ctx.SearchUsers(search)
	if err != nil {
		// TODO: Get error pages working
		c.Redirect(http.StatusMovedPermanently, "/err/500.html")
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
