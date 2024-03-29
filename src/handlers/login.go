/*
 *		LOGIN
 *
 *		All functions pertaining to logging users in/out
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

/// LOGIN

func PgLogin(c echo.Context) error {

	// Check for data
	data := new(struct {
		User  inf.FormUser
		Error inf.UserFormError
	})
	session, err := inf.CookieStore.Get(c.Request(), inf.COOKIE_FORM_DATA)
	if err == nil {
		if e := session.Values["error"]; e != nil {
			if formErr, ok := e.(inf.UserFormError); ok {
				data.Error = formErr
			}
		}
		if u := session.Values["user"]; u != nil {
			if formUsr, ok := u.(inf.FormUser); ok {
				data.User = formUsr
			}
		}
		session.Options.MaxAge = -1
		err = session.Save(c.Request(), c.Response())
		if err != nil {
			fmt.Println("[ERROR] Failed to delete form-data:\n ", err)
			return echo.ErrInternalServerError
		}
	}

	return c.Render(http.StatusOK, "login", data)
}

func LoginUser(c echo.Context) error {

	// Get User data from form
	formUser := new(inf.FormUser)
	err := c.Bind(formUser)
	if err != nil {
		return err
	}

	// Add user to context for if there's an error
	session, err := inf.CookieStore.Get(c.Request(), inf.COOKIE_FORM_DATA)
	if err != nil {
		fmt.Println("[ERROR] Failed to get form-data cookie:\n ", err)
		return echo.ErrInternalServerError
	}
	session.Values["user"] = formUser
	session.Values["error"] = new(inf.UserFormError)

	// Input Validation
	hasError := false
	formError := new(inf.UserFormError)

	// Check all fields were filled
	if formUser.Password == "" {
		hasError = true
		formError.Password = "Password is required"
	}

	wish := inf.GetWishlistClient(wishlistlib.Token{})

	// Get User to log in
	user, err := wish.GetUserByEmail(formUser.Email)
	if err != nil {
		formError.Email = "Please use a valid E-Mail address"
		hasError = true
	}

	// Log in
	err = wish.Authenticate(user.Email, formUser.Password)
	if err != nil {
		hasError = true
		formError.Password = "Email/Password incorrect"
	}

	// Add Token to cookie
	tokenData, err := inf.CookieStore.Get(c.Request(), inf.COOKIE_TOKEN_DATA)
	if err != nil {
		fmt.Println("[ERROR] Failed to get token cookie:\n ", err)
		return echo.ErrInternalServerError
	}
	tokenData.Values[inf.COOKIE_TOKEN_DATA] = wish.Token

	// Set cookie to only expire after a month if the "remember me" option is set
	if formUser.StayLoggedIn == "on" {
		tokenData.Options.MaxAge = inf.COOKIE_TIMEOUT
	}

	err = tokenData.Save(c.Request(), c.Response())
	if err != nil {
		fmt.Println("[ERROR] Failed to save token cookie:\n ", err)
		return echo.ErrInternalServerError
	}

	// Check Errors
	if hasError {
		session.Values["error"] = formError
		err = session.Save(c.Request(), c.Response())
		if err != nil {
			fmt.Println("[ERROR] Failed to save form-data:\n ", err)
			return echo.ErrInternalServerError
		}
		c.Redirect(http.StatusMovedPermanently, "/login")
	}

	return c.Redirect(http.StatusMovedPermanently, "/user/"+user.Email)
}

func Logout(c echo.Context) error {

	// Delete Token Cookie
	userData, err := inf.CookieStore.Get(c.Request(), inf.COOKIE_TOKEN_DATA)
	if err != nil {
		fmt.Println("[ERROR] Failed to get token cookie:\n ", err)
		return echo.ErrInternalServerError
	}
	userData.Options.MaxAge = -1

	err = userData.Save(c.Request(), c.Response())
	if err != nil {
		fmt.Println("[ERROR] Failed to save token cookie:\n ", err)
		return echo.ErrInternalServerError
	}

	return c.Render(http.StatusOK, "status", inf.StatusPageData{
		Colour:          "green",
		MainMessage:     "Successfully logged out!",
		NextPageURL:     "/",
		NextPageMessage: "Back to main page",
	})
}
