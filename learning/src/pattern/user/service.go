package user

var Service IUserService = &userService{}

type IUserService interface {
}

type userService struct {
}
