package userhandler

import (
	"net/http"

	"github.com/labstack/echo"
)

// Create create user handler
func Create(c echo.Context) error {
	return c.String(http.StatusOK, "CreateUserHandler")
}

// SignIn user sign in handler
func SignIn(c echo.Context) error {
	return c.String(http.StatusOK, "user signin handler")
}
