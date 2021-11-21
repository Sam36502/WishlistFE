package inf

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"

	wishlistlib "github.com/Sam36502/WishlistLib-go"
	"github.com/labstack/echo/v4"
)

const (
	PASS_MIN_LEN = 8 // Minimum allowed password length
)

// IsEmailValid checks if the email provided is valid by regex.
// Courtesy of Stackoverflow user "icza"
func IsEmailValid(e string) bool {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return emailRegex.MatchString(e)
}

// Gets and parses the currently logged in user from the cookie
func GetLoggedInUser(c echo.Context) (wishlistlib.User, error) {
	userData, err := CookieStore.Get(c.Request(), COOKIE_USER_DATA)
	if err != nil {
		fmt.Println("[ERROR] Failed to get user-data cookie:\n ", err)
		return wishlistlib.User{}, err
	}

	userStr, exists := userData.Values["user"]
	if !exists {
		return wishlistlib.User{}, errors.New("no user logged in")
	}

	cookieUser, ok := userStr.(CookieUser)
	if !ok {
		fmt.Println("[ERROR] Couldn't convert cookie user. Deleted Cookie.")
		userData.Options.MaxAge = -1
		err = userData.Save(c.Request(), c.Response())
		if err != nil {
			fmt.Println("[ERROR] Failed to delete cookie:\n ", err)
		}
		return wishlistlib.User{}, errors.New("couldn't convert logged in user")
	}

	user := wishlistlib.User{
		ID:    cookieUser.ID,
		Email: cookieUser.Email,
	}
	user.SetPassword(cookieUser.Password)

	return user, nil
}

// Gets all item data based on the "item_id" path parameter
func GetItemFromPath(c echo.Context) (wishlistlib.Item, error) {
	var item_id uint64
	var err error
	if item_id_str := c.Param("item_id"); item_id_str == "" {
		return wishlistlib.Item{}, echo.ErrNotFound
	} else {
		item_id, err = strconv.ParseUint(item_id_str, 10, 64)
		if err != nil {
			return wishlistlib.Item{}, echo.ErrNotFound
		}
	}
	wish := wishlistlib.Context{
		BaseUrl: WISHLIST_BASE_URL,
	}
	item, err := wish.GetItemByID(item_id)
	if err != nil {
		return wishlistlib.Item{}, echo.ErrNotFound
	}

	return item, nil
}

func GetStatusColour(s wishlistlib.Status) string {
	switch s.StatusID {

	default:
		fallthrough
	case 1:
		return "near-white"
	case 2:
		return "gold"
	case 3:
		return "green"

	}
}

// Checks if a given password matches some basic security criteria
// If it matches, this returns an empty string
// otherwise, it returns a string with an error message
// e.g.: "Password must be at least 8 characters"
func CheckPasswordRequirements(password string) string {

	if len(password) < PASS_MIN_LEN {
		return "Password must be at least " + strconv.Itoa(PASS_MIN_LEN) + " characters"
	}

	// TODO: More rigourous constraints

	return ""
}
