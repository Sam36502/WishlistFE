package inf

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ErrorPageContent struct {
	ErrorCode string
	ErrorDesc string
}

func RedirectHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	httpError, ok := err.(*echo.HTTPError)
	if ok {
		code = httpError.Code
	}

	switch code {

	default:
		c.Render(http.StatusOK, "error", ErrorPageContent{
			ErrorCode: fmt.Sprintf("Error %v - %v", httpError.Code, httpError.Message),
			ErrorDesc: "This is quite unusual.\nSend me an E-Mail at sam@aepearce.com with how you got here and the error code, and I'll take a look at it.",
		})
	case 404:
		c.Render(http.StatusOK, "error", ErrorPageContent{
			ErrorCode: "Error 404 - Not Found",
			ErrorDesc: "The page you were looking for couldn't be found.",
		})

	case 500:
		c.Render(http.StatusOK, "error", ErrorPageContent{
			ErrorCode: "Error 500 - Internal Server Error",
			ErrorDesc: "Something went wrong on our end. Please try again later.",
		})

	}
}
