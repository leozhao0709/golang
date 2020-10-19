package errorhandler

import (
	"github.com/labstack/echo"
	"github.com/leozhao0709/golang/bookstore_user_api/app/common"
)

// RestfulHandler ...
func RestfulHandler(err error, c echo.Context) {

	if errorResponse, ok := err.(common.ICommonError); ok {
		c.JSON(errorResponse.StatusCode(), errorResponse)
	} else if he, ok := err.(*echo.HTTPError); ok {
		c.JSON(he.Code, err)
	} else {
		errResponse := common.InternalServerError(err)
		c.JSON(errResponse.StatusCode(), errResponse)
	}
}
