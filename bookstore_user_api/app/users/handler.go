package users

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/leozhao0709/musings/common"
	"github.com/leozhao0709/musings/reflect"
	xerrors "github.com/pkg/errors"
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
	test(c echo.Context) error
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

type testMessage struct {
	Test1 string
	Test2 int
}

func (h *handler) test(c echo.Context) error {
	fmt.Println("...before handle test...")
	// return fmt.Errorf("...error...")
	<-time.After(time.Second * 5)
	// panic(fmt.Errorf("decompress..."))
	fmt.Println("...after handle test...")
	return c.Render(http.StatusOK, "message", testMessage{
		Test1: "test1",
		Test2: 25,
	})
}

func (h *handler) getUser(c echo.Context) error {
	userID := c.Param("user_id")
	log.Info("...", c.Request().Header.Get("Authorization"))
	u, err := h.service.GetUser(c.Request().Context(), userID)
	if err != nil {
		err1 := xerrors.Wrap(err, "789")
		err2 := xerrors.WithMessage(err1, "456")
		log := fmt.Sprintf("%+v", err2)
		fmt.Println(log)
		fmt.Println(xerrors.Is(err2, err1))
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
