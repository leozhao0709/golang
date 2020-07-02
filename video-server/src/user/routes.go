package user

import "github.com/labstack/echo"

// RegisterUserRoute register user route
func RegisterUserRoute(e *echo.Echo) {
	e.POST("/user/signup", CreateHandler)
	e.POST("/user/signin", SignInHandler)
}
