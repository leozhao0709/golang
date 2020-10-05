package user

import "github.com/labstack/echo"

// RegisterRoute register user route
func RegisterRoute(e *echo.Echo) {
	g := e.Group("/user")
	g.POST("/signup", signupHandler)
	g.POST("/login", loginHandler)
	g.DELETE("/delete", deleteHandler)
}
