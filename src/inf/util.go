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
		fmt.Println("[ERROR] Couldn't convert cookie user. Deleted Cookie.")
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
	var jwtClaims TokenJWTClaims
	_, _, err = jwtParser.ParseUnverified(cookieToken.Token, jwtClaims)
	if err != nil {
		return wishlistlib.User{}, wishlistlib.Token{}, errors.New("failed to parse JWT Claims: " + err.Error())
	}

	/*
		jwtClaimsEnc := strings.Split(cookieToken.Token, ".")[1]
		jwtClaimsData, err := base64.StdEncoding.DecodeString(jwtClaimsEnc)
		if err != nil {
			return wishlistlib.User{}, wishlistlib.Token{}, errors.New("failed to decode JWT Claims: " + err.Error())
		}
		var jwtClaims map[string]string
		err = json.Unmarshal(jwtClaimsData, &jwtClaims)
		if err != nil {
			return wishlistlib.User{}, wishlistlib.Token{}, errors.New("failed to parse JWT Claims: " + err.Error())
		}
		email, exists := jwtClaims["email"]
		if !exists {
			return wishlistlib.User{}, wishlistlib.Token{}, errors.New("invalid Token JWT saved")
		}
	*/

	wish := wishlistlib.DefaultWishClient(WISHLIST_BASE_URL)
	user, err := wish.GetUserByEmail(jwtClaims.Email)
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
	wish := wishlistlib.DefaultWishClient(WISHLIST_BASE_URL)
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
