/*
 *		LOGIN
 *
 *		All functions pertaining to logging users in/out
 *
 */

package main

import (
	"fmt"
	"net/http"

	wishlistlib "github.com/Sam36502/WishlistLib-go"
	"github.com/labstack/echo/v4"
)

/// LOGIN

func pgLogin(c echo.Context) error {

	// Check for data
	data := new(struct {
		User  FormUser
		Error FormError
	})
	session, err := CookieStore.Get(c.Request(), COOKIE_FORM_DATA)
	if err == nil {
		if e := session.Values["error"]; e != nil {
			if formErr, ok := e.(FormError); ok {
				data.Error = formErr
			}
		}
		if u := session.Values["user"]; u != nil {
			if formUsr, ok := u.(FormUser); ok {
				data.User = formUsr
			}
		}
		session.Options.MaxAge = -1
		err = session.Save(c.Request(), c.Response())
		if err != nil {
			// TODO: Get error handling sorted
			fmt.Println("[ERROR] Failed to delete form-data:\n ", err)
			return c.Redirect(http.StatusMovedPermanently, "/err/500.html")
		}
	}

	return c.Render(http.StatusOK, "login", data)
}

func loginUser(c echo.Context) error {

	// Get User data from form
	formUser := new(FormUser)
	err := c.Bind(formUser)
	if err != nil {
		return err
	}

	// Add user to context for if there's an error
	session, err := CookieStore.Get(c.Request(), COOKIE_FORM_DATA)
	if err != nil {
		// TODO: Get error pages working
		fmt.Println("[ERROR] Failed to get form-data cookie:\n ", err)
		c.Redirect(http.StatusMovedPermanently, "/err/500.html")
	}
	session.Values["user"] = formUser
	session.Values["error"] = new(FormError)

	// Input Validation
	hasError := false
	formError := new(FormError)

	// Check all fields were filled
	if formUser.Password == "" {
		hasError = true
		formError.Password = "Password & confirmation are required"
	}

	// Check Email is Valid
	if !isEmailValid(formUser.Email) {
		formError.Email = "Please use a valid E-Mail address"
		hasError = true
	}

	// Log user in
	user := wishlistlib.User{
		Email: formUser.Email,
	}
	user.SetPassword(formUser.Password)

	wish := wishlistlib.Context{
		BaseUrl: WISHLIST_BASE_URL,
	}
	wish.SetAuthenticatedUser(user)

	// Check Login
	err = wish.CheckCredentials()
	if err != nil {
		hasError = true
		formError.Password = "Email/Password incorrect"
	}

	// Add user data to cookie
	userData, err := CookieStore.Get(c.Request(), COOKIE_USER_DATA)
	if err != nil {
		// TODO: Get error pages working
		fmt.Println("[ERROR] Failed to get user-data cookie:\n ", err)
		c.Redirect(http.StatusMovedPermanently, "/err/500.html")
	}
	userData.Values["user"] = CookieUser{
		Email:    user.Email,
		Password: formUser.Password,
	}

	// Set Cookie to expire if "keep me logged in" not checked
	if !(formUser.StayLoggedIn == "on") {
		userData.Options.MaxAge = COOKIE_TIMEOUT
	}

	err = userData.Save(c.Request(), c.Response())
	if err != nil {
		// TODO: Get error pages working
		fmt.Println("[ERROR] Failed to save user-data:\n ", err)
		c.Redirect(http.StatusMovedPermanently, "/err/500.html")
	}

	if hasError {
		session.Values["error"] = formError
		err = session.Save(c.Request(), c.Response())
		if err != nil {
			// TODO: Get error pages working
			fmt.Println("[ERROR] Failed to save form-data:\n ", err)
			c.Redirect(http.StatusMovedPermanently, "/err/500.html")
		}
		c.Redirect(http.StatusMovedPermanently, "/login")
	}

	return c.Redirect(http.StatusMovedPermanently, "/user/"+user.Email)
}
