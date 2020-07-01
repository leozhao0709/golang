package routes

import (
	"github.com/labstack/echo"
	userhandler "github.com/leozhao0709/golang/video-server/src/api/handler/user"
)

// RegisterUserRoute register user route
func RegisterUserRoute(e *echo.Echo) {
	e.POST("/user/signup", userhandler.Create)
	e.POST("/user/signin", userhandler.SignIn)
}
