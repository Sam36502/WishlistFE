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
		Error inf.FormError
	})
	session, err := inf.CookieStore.Get(c.Request(), inf.COOKIE_FORM_DATA)
	if err == nil {
		if e := session.Values["error"]; e != nil {
			if formErr, ok := e.(inf.FormError); ok {
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
	session.Values["error"] = new(inf.FormError)

	// Input Validation
	hasError := false
	formError := new(inf.FormError)

	// Check all fields were filled
	if formUser.Password == "" {
		hasError = true
		formError.Password = "Password & confirmation are required"
	}

	// Check Email is Valid
	if !inf.IsEmailValid(formUser.Email) {
		formError.Email = "Please use a valid E-Mail address"
		hasError = true
	}

	// Log user in
	user := wishlistlib.User{
		Email: formUser.Email,
	}
	user.SetPassword(formUser.Password)

	wish := wishlistlib.Context{
		BaseUrl: inf.WISHLIST_BASE_URL,
	}
	wish.SetAuthenticatedUser(user)

	// Check Login
	err = wish.CheckCredentials()
	if err != nil {
		hasError = true
		formError.Password = "Email/Password incorrect"
	}

	// Add user data to cookie
	userData, err := inf.CookieStore.Get(c.Request(), inf.COOKIE_USER_DATA)
	if err != nil {
		fmt.Println("[ERROR] Failed to get user-data cookie:\n ", err)
		return echo.ErrInternalServerError
	}
	userData.Values["user"] = inf.CookieUser{
		Email:    user.Email,
		Password: formUser.Password,
	}

	// Set Cookie to expire if "keep me logged in" not checked
	if !(formUser.StayLoggedIn == "on") {
		userData.Options.MaxAge = inf.COOKIE_TIMEOUT
	}

	err = userData.Save(c.Request(), c.Response())
	if err != nil {
		fmt.Println("[ERROR] Failed to save user-data:\n ", err)
		return echo.ErrInternalServerError
	}

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
