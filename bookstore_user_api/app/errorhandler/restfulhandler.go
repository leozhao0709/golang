package errorhandler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/leozhao0709/golang/bookstore_user_api/app/response"
)

// RestfulHandler ...
func RestfulHandler(err error, c echo.Context) {

	if errorResponse, ok := err.(response.IErrorResponse); ok {
		c.JSON(errorResponse.StatusCode(), errorResponse)
	} else if he, ok := err.(*echo.HTTPError); ok {
		c.JSON(he.Code, err)
	} else {
		c.JSON(http.StatusInternalServerError, err)
	}
}
