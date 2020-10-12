package users

var (
	// Service user service
	Service IUserService = &userService{}
)

// IUserService User service interface
type IUserService interface {
}

type userService struct {
}
