package user

import "github.com/labstack/echo"

// RegisterRoute register user route
func RegisterRoute(e *echo.Echo) {
	e.POST("/user", createUser)
	e.POST("/user/:username", loginUser)
}
