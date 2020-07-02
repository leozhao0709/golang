package user

import (
	"net/http"

	"github.com/labstack/echo"
)

// Create create user handler
func CreateHandler(c echo.Context) error {
	return c.String(http.StatusOK, "CreateUserHandler")
}

// SignIn user sign in handler
func SignInHandler(c echo.Context) error {
	return c.String(http.StatusOK, "user signin handler")
}
