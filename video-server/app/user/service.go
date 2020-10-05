package user

import (
	"sync"

	"github.com/labstack/echo"
	"github.com/leozhao0709/golang/video-server/app/db"
	"github.com/leozhao0709/golang/video-server/ent"
	"github.com/leozhao0709/golang/video-server/ent/user"
)

var once sync.Once
var srv *service

// Service user service interface
type Service interface {
	CreateUser(username string, pwd string, ctx echo.Context) (*ent.User, error)
	LoginUser(username string, pwd string, ctx echo.Context) (*ent.User, error)
	DeleteUser(username string, pwd string, ctx echo.Context) error
}

type service struct{}

// GetService create new service
func GetService() Service {
	once.Do(func() {
		srv = &service{}
	})

	return srv
}

func (s *service) CreateUser(username string, pwd string, ctx echo.Context) (*ent.User, error) {
	client, err := db.GetEntClient()
	if err != nil {
		return nil, err
	}

	return client.User.
		Create().
		SetUsername(username).
		SetPassword(pwd).
		Save(ctx.Request().Context())
}

func (s *service) LoginUser(username string, pwd string, ctx echo.Context) (*ent.User, error) {
	client, err := db.GetEntClient()
	if err != nil {
		return nil, err
	}

	return client.User.
		Query().
		Where(user.UsernameEQ(username)).
		Where(user.PasswordEQ(pwd)).
		First(ctx.Request().Context())
}

func (s *service) DeleteUser(username string, pwd string, ctx echo.Context) error {
	client, err := db.GetEntClient()
	if err != nil {
		return err
	}

	_, err = client.User.
		Delete().
		Where(user.UsernameEQ(username)).
		Where(user.PasswordEQ(pwd)).Exec(ctx.Request().Context())

	return err
}
