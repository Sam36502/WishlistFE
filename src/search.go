package main

import (
	"net/http"
	"strconv"

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
	if !c.QueryParams().Has(SEARCH_QUERY_PARAM) {
		return c.Render(http.StatusOK, "search", searchResultData{
			Search:       "",
			ResultString: "Enter a name above to find friends.",
			Results:      nil,
		})
	}
	search := c.QueryParam(SEARCH_QUERY_PARAM)

	// Search Users
	ctx := wishlist.DefaultContext()
	users, err := ctx.SearchUsers(search)
	if err != nil {
		c.Redirect(308, "/err/500.html")
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
