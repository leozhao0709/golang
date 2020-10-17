package users

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/leozhao0709/golang/bookstore_user_api/app/db"
	"github.com/leozhao0709/golang/bookstore_user_api/app/response"
	"github.com/leozhao0709/golang/bookstore_user_api/ent/user"
	"github.com/leozhao0709/musings/reflect"
	"golang.org/x/crypto/bcrypt"
)

var (
	// Service user service
	Service IService = &service{}
)

// IService User service interface
type IService interface {
	CreateUser(echo.Context, User) (response.IResponse, error)
	// GetUser(string) (User, error)
	// UpdateUser(bool, User) (*User, error)
	// DeleteUser(int64) error
	// SearchUser(string) (User, error)
}

type service struct{}

func (srv *service) CreateUser(c echo.Context, u User) (response.IResponse, error) {

	if err := u.Validate(); err != nil {
		return nil, response.BadRequest(err)
	}

	userCount, err := db.GetEntClient().User.Query().Where(user.EmailEQ(u.Email)).Count(c.Request().Context())
	if err != nil {
		return nil, response.InternalServerError(err)
	}

	if userCount > 0 {
		return nil, response.BadRequest(fmt.Errorf("user with email %s already exist", u.Email))
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		return nil, response.InternalServerError(err)
	}

	entUser, err := db.GetEntClient().User.Create().SetEmail(u.Email).SetPassword(string(hashedPassword)).SetFirstName(u.FirstName).SetLastName(u.LastName).Save(c.Request().Context())

	if err != nil {
		return nil, response.InternalServerError(err)
	}

	publicUser := &PublicUser{}
	reflect.CopyProperties(entUser, publicUser)

	return response.SuccessResponse(publicUser), nil
}

func (srv *service) GetUser(c echo.Context, u User) (response.IResponse, error) {
	if err := u.Validate(); err != nil {
		return nil, response.BadRequest(err)
	}

	entUser, err := db.GetEntClient().User.Query().Where(user.EmailEQ(u.Email)).Only(c.Request().Context())
	if err != nil {
		return nil, response.InternalServerError(err)
	}

	user := &User{}
	reflect.CopyProperties(entUser, user)

	return response.SuccessResponse(user), nil
}

// func (srv *service) UpdateUser(bool, User) (*User, error) {

// }

// func (srv *service) DeleteUser(int64) error {

// }

// func (srv *service) SearchUser(string) (User, error) {

// }
