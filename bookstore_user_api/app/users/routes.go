package users

import (
	"github.com/labstack/echo"
)

// RegisterRoute register user route
func RegisterRoute(e *echo.Echo) {

	handler := InjectHandler()

	g := e.Group("/user")
	g.POST("/create", handler.createUser)
	g.GET("/:user_id", handler.getUser)
	g.DELETE("/:user_id", handler.deleteUser)
	g.PATCH("/:user_id", handler.updateUser)
}
