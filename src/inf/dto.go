/*
 *		D.T.Os (Data Transfer Objects)
 *
 *		All structs for packaging data to pass
 *		through to the pages
 *
 */

package inf

type FormUser struct {
	Name            string `form:"name"`
	Email           string `form:"email"`
	Password        string `form:"password"`
	PasswordConfirm string `form:"pass_conf"`
	StayLoggedIn    string `form:"stay-logged-in"`
}

type FormError struct {
	Name     string
	Email    string
	Password string
}
