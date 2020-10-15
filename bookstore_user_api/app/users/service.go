package users

import "github.com/leozhao0709/golang/bookstore_user_api/app/common"

var (
	// Service user service
	Service IService = &service{}
)

// IService User service interface
type IService interface {
	CreateUser(User) (common.IResponse, common.IErrorResponse)
	// GetUser(int64) (User, error)
	// UpdateUser(bool, User) (*User, error)
	// DeleteUser(int64) error
	// SearchUser(string) (User, error)
}

type service struct{}

func (srv *service) CreateUser(u User) (common.IResponse, common.IErrorResponse) {
	return common.SuccessResponse(100, u), nil
}

// func (srv *service) GetUser(int64) (User, error) {

// }

// func (srv *service) UpdateUser(bool, User) (*User, error) {

// }

// func (srv *service) DeleteUser(int64) error {

// }

// func (srv *service) SearchUser(string) (User, error) {

// }
