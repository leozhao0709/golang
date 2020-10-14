package user

var Service IUserService = &userService{}

type IUserService interface {
	GetAge() int
}

type userService struct {
}

func (service userService) GetAge() int {
	return 29
}
