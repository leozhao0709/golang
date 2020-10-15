package users

import (
	"github.com/labstack/echo"
)

// func getUserHandler(c echo.Context) error {
// 	return c.String(http.StatusOK, "Get user")
// }

func createUserHandler(c echo.Context) error {
	u := &User{}
	if err := c.Bind(u); err != nil {
		return err
	}

	response, err := Service.CreateUser(*u)
	if err != nil {
		return err
	}

	return c.JSON(response.StatusCode(), response)
}
