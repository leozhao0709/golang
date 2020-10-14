package users

var (
	// Service user service
	Service IUserService = &userService{}
)

// IUserService User service interface
type IUserService interface {
	GetUser(int64) (User, error)
	CreateUser(User) (User, error)
	UpdateUser(bool, User) (*User, error)
	DeleteUser(int64) error
	SearchUser(string) (User, error)
}

type userService struct {
}

func (service userService) GetUser(int64) (User, error) {

}

func (service userService) CreateUser(User) (User, error) {

}

func (service userService) UpdateUser(bool, User) (*User, error) {

}

func (service userService) DeleteUser(int64) error {

}

func (service userService) SearchUser(string) (User, error) {

}
