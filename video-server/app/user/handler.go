package user

import (
	"net/http"

	"github.com/labstack/echo"
)

func signupHandler(c echo.Context) error {
	userCredential := &Credential{}
	if err := c.Bind(userCredential); err != nil {
		return err
	}

	userService := GetService()
	user, err := userService.CreateUser(userCredential.Username, userCredential.Password, c)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func loginHandler(c echo.Context) error {
	userCredential := &Credential{}
	if err := c.Bind(userCredential); err != nil {
		return err
	}

	userService := GetService()
	user, err := userService.LoginUser(userCredential.Username, userCredential.Password, c)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

func deleteHandler(c echo.Context) error {
	userCredential := &Credential{}
	if err := c.Bind(userCredential); err != nil {
		return err
	}

	userService := GetService()
	err := userService.DeleteUser(userCredential.Username, userCredential.Password, c)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]bool{
		"success": true,
	})
}
