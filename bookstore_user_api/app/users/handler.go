package users

import (
	"net/http"

	"github.com/labstack/echo"
)

func getUserHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Get user")
}
