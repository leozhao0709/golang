package users

import (
	"sync"

	"github.com/labstack/echo"
	"github.com/leozhao0709/musings/common"
	"github.com/leozhao0709/musings/reflect"
)

var (
	h     *handler
	hOnce sync.Once
)

// IHandler ...
type IHandler interface {
	getUser(c echo.Context) error
	createUser(c echo.Context) error
	deleteUser(c echo.Context) error
	updateUser(c echo.Context) error
}

type handler struct {
	service IService
}

// GetHandler ...
func GetHandler(service IService) IHandler {
	hOnce.Do(func() {
		h = &handler{
			service: service,
		}
	})
	return h
}

func (h *handler) getUser(c echo.Context) error {
	userID := c.Param("user_id")

	u, err := h.service.GetUser(c.Request().Context(), userID)
	if err != nil {
		return err
	}

	response := createResponseUser(c, u)

	return c.JSON(response.StatusCode(), response.Data())
}

func (h *handler) createUser(c echo.Context) error {
	u := &User{}
	if err := c.Bind(u); err != nil {
		return err
	}

	u, err := h.service.CreateUser(c.Request().Context(), u)
	if err != nil {
		return err
	}

	response := createResponseUser(c, u)

	return c.JSON(response.StatusCode(), response.Data())
}

func (h *handler) deleteUser(c echo.Context) error {
	userID := c.Param("user_id")

	num, err := h.service.DeleteUser(c.Request().Context(), userID)
	if err != nil {
		return err
	}

	response := common.SuccessResponse(map[string]int{"deleteNum": num})

	return c.JSON(response.StatusCode(), response.Data())
}

func (h *handler) updateUser(c echo.Context) error {
	userID := c.Param("user_id")

	u := &User{}
	if err := c.Bind(u); err != nil {
		return err
	}

	num, err := h.service.UpdateUser(c.Request().Context(), userID, u)

	if err != nil {
		return err
	}

	response := common.SuccessResponse(map[string]int{"updateNum": num})

	return c.JSON(response.StatusCode(), response.Data())
}

// private
func createResponseUser(c echo.Context, user *User) common.IResponse {

	var isPublic = c.Request().Header.Get("X-Public") == "true"

	var respUser interface{}
	if isPublic {
		respUser = &PublicUser{}
	} else {
		respUser = &PrivateUser{}
	}

	reflect.CopyProperties(user, respUser)

	return common.SuccessResponse(respUser)
}
