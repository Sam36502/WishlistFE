package inf

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"

	wishlistlib "github.com/Sam36502/WishlistLib-go"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

const (
	PASS_MIN_LEN = 8 // Minimum allowed password length

	STATUS_AVAILABLE = 1
	STATUS_RESERVED  = 2
	STATUS_RECEIVED  = 3
)

// IsEmailValid checks if the email provided is valid by regex.
// Courtesy of Stackoverflow user "icza"
func IsEmailValid(e string) bool {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return emailRegex.MatchString(e)
}

// Gets and parses the currently logged in user and token from the cookie
func GetLoggedInUser(c echo.Context) (wishlistlib.User, wishlistlib.Token, error) {
	tokenData, err := CookieStore.Get(c.Request(), COOKIE_TOKEN_DATA)
	if err != nil {
		fmt.Println("[ERROR] Failed to get token cookie:\n ", err)
		return wishlistlib.User{}, wishlistlib.Token{}, err
	}

	tokenInterface, exists := tokenData.Values[COOKIE_TOKEN_DATA]
	if !exists {
		return wishlistlib.User{}, wishlistlib.Token{}, errors.New("no user logged in")
	}

	cookieToken, ok := tokenInterface.(wishlistlib.Token)
	if !ok {
		fmt.Println("[ERROR] Couldn't convert cookie user. Deleting Cookie...")
		tokenData.Options.MaxAge = -1
		err = tokenData.Save(c.Request(), c.Response())
		if err != nil {
			fmt.Println("[ERROR] Failed to delete cookie:\n ", err)
		}
		return wishlistlib.User{}, wishlistlib.Token{}, errors.New("failed to convert cookie token")
	}

	// Parse email from JWT
	jwtParser := jwt.Parser{
		ValidMethods:         []string{},
		UseJSONNumber:        false,
		SkipClaimsValidation: true,
	}
	token, _, err := jwtParser.ParseUnverified(cookieToken.Token, jwt.MapClaims{})
	if err != nil {
		return wishlistlib.User{}, wishlistlib.Token{}, errors.New("failed to parse JWT Claims: " + err.Error())
	}
	jwtClaims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return wishlistlib.User{}, wishlistlib.Token{}, errors.New("failed to convert cookie claims to token claims")
	}
	emailIfc, exists := jwtClaims["email"]
	if !exists {
		return wishlistlib.User{}, wishlistlib.Token{}, errors.New("invalid JWT Claims parsed; no 'email' field present")
	}
	email, ok := emailIfc.(string)
	if !ok {
		return wishlistlib.User{}, wishlistlib.Token{}, errors.New("invalid JWT Claims parsed; email is not a string")
	}

	wish := GetWishlistClient(wishlistlib.Token{})
	user, err := wish.GetUserByEmail(email)
	if err != nil {
		return wishlistlib.User{}, wishlistlib.Token{}, errors.New("failed to retrieve logged-in user from API: " + err.Error())
	}

	return user, cookieToken, nil
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
	wish := GetWishlistClient(wishlistlib.Token{})
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
