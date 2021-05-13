package users

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/leozhao0709/musings/common"
)

// RegisterRoute register user route
func RegisterRoute(e *echo.Echo) {

	handler := InjectHandler()

	g := e.Group("/user")

	m1 := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			fmt.Println(".....m1 before...")
			err := next(c)
			fmt.Println("...m1 err...", err)
			fmt.Println("...m1 after...")
			if err != nil {
				errResponse := common.InternalServerError(fmt.Errorf("inner error"))
				return c.JSON(errResponse.StatusCode(), errResponse)
			}
			return nil
		}
	}

	m2 := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			fmt.Println(".....m2 before...")
			err := next(c)
			fmt.Println("...m2 err...", err)
			fmt.Println("...m2 after...")

			return err
		}
	}

	g.Use(m1, m2)

	g.GET("/test", handler.test)
	g.POST("/create", handler.createUser)
	g.GET("/:user_id", handler.getUser)
	g.DELETE("/:user_id", handler.deleteUser)
	g.PATCH("/:user_id", handler.updateUser)
}
