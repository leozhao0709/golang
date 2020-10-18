package users

import "github.com/labstack/echo"

// RegisterRoute register user route
func RegisterRoute(e *echo.Echo) {
	g := e.Group("/user")
	g.POST("/create", createUserHandler)
	g.GET("/:user_id", getUserHandler)
	g.DELETE("/:user_id", deleteUserHandler)
	g.PATCH("/:user_id", updateUserHandler)
}
