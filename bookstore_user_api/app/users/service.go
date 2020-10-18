package users

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/leozhao0709/golang/bookstore_user_api/app/db"
	"github.com/leozhao0709/golang/bookstore_user_api/app/errcode"
	"github.com/leozhao0709/golang/bookstore_user_api/app/response"
	"github.com/leozhao0709/golang/bookstore_user_api/ent"
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
	CreateUser(c echo.Context, u *User, isPublic bool) (response.IResponse, error)
	GetUser(c echo.Context, userID string, isPublic bool) (response.IResponse, error)
	DeleteUser(c echo.Context, userID string, isPublic bool) (response.IResponse, error)
	UpdateUser(c echo.Context, userID string, u *User, isPublic bool) (response.IResponse, error)
	// SearchUser(c echo.Context, enable bool) (response.IResponse, error)
}

type service struct{}

func (srv *service) CreateUser(c echo.Context, u *User, isPublic bool) (response.IResponse, error) {

	if err := createValidate(*u); err != nil {
		return nil, response.BadRequest(errcode.RequestValidationErr, err)
	}

	exist, err := db.GetEntClient().User.Query().Where(user.EmailEQ(u.Email)).Limit(1).Exist(c.Request().Context())
	if err != nil {
		return nil, response.InternalServerError(err)
	}

	if exist {
		return nil, response.BadRequest(userExistErrCode, fmt.Errorf("user with email %s already exist", u.Email))
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		return nil, response.InternalServerError(err)
	}

	entUser, err := db.GetEntClient().User.Create().SetEmail(u.Email).SetPassword(string(hashedPassword)).SetFirstName(u.FirstName).SetLastName(u.LastName).Save(c.Request().Context())

	if err != nil {
		return nil, response.InternalServerError(err)
	}

	var respUser = createResponseUser(entUser, isPublic)

	return response.SuccessResponse(respUser), nil
}

func (srv *service) GetUser(c echo.Context, userID string, isPublic bool) (response.IResponse, error) {
	uid, err := uuid.Parse(userID)
	if err != nil {
		return nil, response.BadRequest(BadUserIDCode, err)
	}

	entUsers, err := db.GetEntClient().User.Query().Where(user.UserIDEQ(uid)).Limit(1).All(c.Request().Context())
	if err != nil {
		return nil, response.InternalServerError(err)
	}

	if len(entUsers) == 0 {
		return nil, response.BadRequest(userNotFoundErrCode, fmt.Errorf("user not found"))
	}

	respUser := createResponseUser(entUsers[0], isPublic)

	return response.SuccessResponse(respUser), nil
}

func (srv *service) UpdateUser(c echo.Context, userID string, u *User, isPublic bool) (response.IResponse, error) {
	if err := updateValidate(*u); err != nil {
		return nil, response.BadRequest(errcode.RequestValidationErr, err)
	}

	uid, err := uuid.Parse(userID)
	if err != nil {
		return nil, response.BadRequest(BadUserIDCode, err)
	}

	entBuilder := db.GetEntClient().User.Update().Where(user.UserIDEQ(uid))

	if u.Email != "" {
		entBuilder = entBuilder.SetEmail(u.Email)
	}

	if u.Password != "" {
		entBuilder = entBuilder.SetPassword(u.Password)
	}

	if u.FirstName != "" {
		entBuilder = entBuilder.SetFirstName(u.FirstName)
	}

	if u.LastName != "" {
		entBuilder = entBuilder.SetLastName(u.LastName)
	}

	if u.Status != "" {
		entBuilder = entBuilder.SetStatus(u.Status)
	}

	updateCount, err := entBuilder.Save(c.Request().Context())

	if err != nil {
		return nil, response.InternalServerError(err)
	}

	return response.SuccessResponse(map[string]int{"updateNum": updateCount}), nil
}

func (srv *service) DeleteUser(c echo.Context, userID string, isPublic bool) (response.IResponse, error) {

	uid, err := uuid.Parse(userID)
	if err != nil {
		return nil, response.BadRequest(BadUserIDCode, err)
	}

	count, err := db.GetEntClient().User.Delete().Where(user.UserIDEQ(uid)).Exec(c.Request().Context())

	if err != nil {
		return nil, response.InternalServerError(err)
	}

	return response.SuccessResponse(map[string]int{"deleteNum": count}), nil
}

// func (srv *service) SearchUser(string) (User, error) {

// }

// private
func createResponseUser(entUser *ent.User, isPublic bool) interface{} {
	var respUser interface{}
	if isPublic {
		respUser = &PublicUser{}
	} else {
		respUser = &PrivateUser{}
	}

	reflect.CopyProperties(entUser, respUser)

	return respUser
}

func createValidate(u User) error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Status, validation.In("enable", "disable")),
		validation.Field(&u.Password, validation.Required),
	)
}

func updateValidate(u User) error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Email, is.Email),
		validation.Field(&u.Status, validation.In("enable", "disable")),
	)
}
