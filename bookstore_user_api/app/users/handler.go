package users

import (
	"github.com/labstack/echo"
)

func getUserHandler(c echo.Context) error {
	userID := c.Param("user_id")

	response, err := Service.GetUser(c, userID, isPublicUser(c))
	if err != nil {
		return err
	}

	return c.JSON(response.StatusCode(), response.Data())
}

func createUserHandler(c echo.Context) error {
	u := &User{}
	if err := c.Bind(u); err != nil {
		return err
	}

	response, err := Service.CreateUser(c, u, isPublicUser(c))
	if err != nil {
		return err
	}

	return c.JSON(response.StatusCode(), response.Data())
}

func deleteUserHandler(c echo.Context) error {
	userID := c.Param("user_id")

	response, err := Service.DeleteUser(c, userID, isPublicUser(c))
	if err != nil {
		return err
	}

	return c.JSON(response.StatusCode(), response.Data())
}

func updateUserHandler(c echo.Context) error {
	userID := c.Param("user_id")

	u := &User{}
	if err := c.Bind(u); err != nil {
		return err
	}

	response, err := Service.UpdateUser(c, userID, u, isPublicUser(c))

	if err != nil {
		return err
	}

	return c.JSON(response.StatusCode(), response.Data())
}

// private
func isPublicUser(c echo.Context) bool {
	return c.Request().Header.Get("X-Public") == "true"
}
