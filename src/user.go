package main

import (
	"fmt"
	"net/http"

	wishlistlib "github.com/Sam36502/WishlistLib-go"
	"github.com/labstack/echo/v4"
)

type FormUser struct {
	Name            string `form:"name"`
	Email           string `form:"email"`
	Password        string `form:"password"`
	PasswordConfirm string `form:"pass_conf"`
}

type FormError struct {
	Name     string
	Email    string
	Password string
}

/// REGISTER

func pgRegister(c echo.Context) error {

	// Check for data
	data := new(struct {
		User  FormUser
		Error FormError
	})
	session, err := CookieStore.Get(c.Request(), "register-form")
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
	}
	session.Options.MaxAge = -1
	err = session.Save(c.Request(), c.Response())
	if err != nil {
		// TODO: Get error handling sorted
		fmt.Println("[ERROR] Failed to save session:\n ", err)
		return c.Redirect(http.StatusMovedPermanently, "/err/500.html")
	}

	return c.Render(http.StatusOK, "register", data)
}

func registerUser(c echo.Context) error {

	// Get User data from form
	formUser := new(FormUser)
	err := c.Bind(formUser)
	if err != nil {
		return err
	}

	// Add user to context for if there's an error
	session, err := CookieStore.Get(c.Request(), "register-form")
	if err != nil {
		// TODO: Get error pages working
		fmt.Println("[ERROR] Failed to get Session:\n ", err)
		c.Redirect(http.StatusMovedPermanently, "/err/500.html")
	}
	session.Values["user"] = formUser
	session.Values["error"] = new(FormError)

	// Input Validation
	hasError := false
	formError := new(FormError)

	// Check all fields were filled
	if formUser.Name == "" {
		hasError = true
		formError.Name = "Name is required"
	}
	if formUser.Password == "" || formUser.PasswordConfirm == "" {
		hasError = true
		formError.Password = "Password & confirmation are required"
	}

	// Check Email is Valid
	if !isEmailValid(formUser.Email) {
		formError.Email = "Please use a valid E-Mail address"
		hasError = true
	}

	// Check Password aligns with confirm field
	if formUser.Password != formUser.PasswordConfirm {
		formError.Password = "Passwords don't match."
		hasError = true
	}

	var user wishlistlib.User
	if !hasError {
		// Register the User
		user = wishlistlib.User{
			Name:  formUser.Name,
			Email: formUser.Email,
		}
		user.SetPassword(formUser.Password)
		user, err = Wishlist.AddNewUser(user)
		if err != nil {
			hasError = true
			if _, ok := err.(wishlistlib.EmailExistsError); ok {
				formError.Email = "This Email is already in use"
			}
		}
	}

	if hasError {
		session.Values["error"] = formError
		session.Save(c.Request(), c.Response())
		c.Redirect(http.StatusMovedPermanently, "/register")
	}

	return c.Redirect(http.StatusMovedPermanently, "/user/"+user.Email)
}

/// LOGIN

func pgLogin(c echo.Context) error {

	// Check for data
	data := new(struct {
		User  FormUser
		Error FormError
	})
	session, err := CookieStore.Get(c.Request(), "register-form")
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
	}
	session.Options.MaxAge = -1
	err = session.Save(c.Request(), c.Response())
	if err != nil {
		// TODO: Get error handling sorted
		fmt.Println("[ERROR] Failed to save session:\n ", err)
		return c.Redirect(http.StatusMovedPermanently, "/err/500.html")
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
	session, err := CookieStore.Get(c.Request(), "register-form")
	if err != nil {
		// TODO: Get error pages working
		fmt.Println("[ERROR] Failed to get Session:\n ", err)
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

	// Register the User
	user := wishlistlib.User{
		Email: formUser.Email,
	}
	user.SetPassword(formUser.Password)
	Wishlist.SetAuthenticatedUser(user)

	// Check Login
	err = Wishlist.CheckCredentials()
	if err != nil {
		hasError = true
		formError.Password = "Email/Password incorrect"
	}

	if hasError {
		session.Values["error"] = formError
		session.Save(c.Request(), c.Response())
		c.Redirect(http.StatusMovedPermanently, "/login")
	}

	return c.Redirect(http.StatusMovedPermanently, "/user/"+user.Email)
}

/// USER'S WISHLIST

func pgUserList(c echo.Context) error {

	// Get all the user's items
	email := c.Param("email")
	if email == "" {
		return c.Redirect(http.StatusPermanentRedirect, "/home")
	}

	user, err := Wishlist.GetUserByEmail(email)
	if err != nil {
		return c.Redirect(http.StatusPermanentRedirect, "/home")
	}
	items, err := Wishlist.GetAllItems(user)
	if err != nil {
		fmt.Println("[ERROR] Failed to retrieve User's items:\n ", err)
		return c.Redirect(http.StatusPermanentRedirect, "/err/500.html")
	}

	return c.Render(http.StatusOK, "userlist", struct {
		User  wishlistlib.User
		Items []wishlistlib.Item
	}{
		User:  user,
		Items: items,
	})
}
