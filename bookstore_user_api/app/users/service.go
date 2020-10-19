package users

import (
	"context"
	"fmt"
	"sync"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
	"github.com/leozhao0709/musings/common"
	"github.com/leozhao0709/musings/reflect"
	"golang.org/x/crypto/bcrypt"
)

var (
	s     *service
	sOnce sync.Once
)

// GetService ...
func GetService(repositroy IRepository) IService {
	sOnce.Do(func() {
		s = &service{
			repositroy: repositroy,
		}
	})

	return s
}

// IService User service interface
type IService interface {
	CreateUser(ctx context.Context, u *User) (*User, error)
	GetUser(ctx context.Context, userID string) (*User, error)
	DeleteUser(ctx context.Context, userID string) (int, error)
	UpdateUser(ctx context.Context, userID string, u *User) (int, error)
	// SearchUser(ctx context.Context, enable bool) (response.IResponse, error)
}

type service struct {
	repositroy IRepository
}

func (s *service) CreateUser(ctx context.Context, u *User) (*User, error) {

	if err := createValidate(*u); err != nil {
		return nil, common.BadRequestError(common.RequestValidationErrCode, err)
	}

	exist, err := s.repositroy.IsUserExist(ctx, u)
	if err != nil {
		return nil, common.InternalServerError(err)
	}

	if exist {
		return nil, common.BadRequestError(userExistErrCode, fmt.Errorf("user with email %s already exist", u.Email))
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		return nil, common.InternalServerError(err)
	}

	u.Password = string(hashedPassword)

	entUser, err := s.repositroy.CreateUser(ctx, u)

	if err != nil {
		return nil, common.InternalServerError(err)
	}

	var user = &User{}
	reflect.CopyProperties(entUser, user)

	return user, nil
}

func (s *service) GetUser(ctx context.Context, userID string) (*User, error) {
	uid, err := uuid.Parse(userID)
	if err != nil {
		return nil, common.BadRequestError(BadUserIDCode, err)
	}

	entUser, err := s.repositroy.GetUser(ctx, uid)
	if err != nil {
		return nil, common.InternalServerError(err)
	}

	if entUser == nil {
		return nil, common.BadRequestError(userNotFoundErrCode, fmt.Errorf("user not found"))
	}

	var user = &User{}
	reflect.CopyProperties(entUser, user)

	return user, nil
}

func (s *service) UpdateUser(ctx context.Context, userID string, u *User) (int, error) {
	if err := updateValidate(*u); err != nil {
		return 0, common.BadRequestError(common.RequestValidationErrCode, err)
	}

	uid, err := uuid.Parse(userID)
	if err != nil {
		return 0, common.BadRequestError(BadUserIDCode, err)
	}

	if u.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
		if err != nil {
			return 0, common.InternalServerError(err)
		}

		u.Password = string(hashedPassword)
	}

	return s.repositroy.UpdateUser(ctx, uid, u)
}

func (s *service) DeleteUser(ctx context.Context, userID string) (int, error) {

	uid, err := uuid.Parse(userID)
	if err != nil {
		return 0, common.BadRequestError(BadUserIDCode, err)
	}

	return s.repositroy.DeleteUser(ctx, uid)
}

// func (s *service) SearchUser(string) (User, error) {

// }

// private

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
