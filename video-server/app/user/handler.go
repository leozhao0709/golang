package user

import (
	"net/http"

	"github.com/labstack/echo"
)

func createUser(c echo.Context) error {
	return c.String(http.StatusOK, "create user")
}

func loginUser(c echo.Context) error {
	return c.String(http.StatusOK, c.Param("username"))
}
