package users

import (
	"context"
	"sync"

	"github.com/google/uuid"
	"github.com/leozhao0709/golang/bookstore_user_api/app/common"
	"github.com/leozhao0709/golang/bookstore_user_api/ent"
	"github.com/leozhao0709/golang/bookstore_user_api/ent/user"
)

var (
	r     *repositroy
	rOnce sync.Once
)

// GetRepository ...
func GetRepository(client *ent.UserClient) IRepository {
	rOnce.Do(func() {
		r = &repositroy{
			client: client,
		}
	})

	return r
}

// IRepository ...
type IRepository interface {
	IsUserExist(ctx context.Context, u *User) (bool, error)
	CreateUser(ctx context.Context, u *User) (*ent.User, error)
	GetUser(ctx context.Context, userID uuid.UUID) (*ent.User, error)
	UpdateUser(ctx context.Context, userID uuid.UUID, u *User) (int, error)
	DeleteUser(ctx context.Context, userID uuid.UUID) (int, error)
}

type repositroy struct {
	client *ent.UserClient
}

func (r *repositroy) IsUserExist(ctx context.Context, u *User) (bool, error) {
	return r.client.Query().Where(user.EmailEQ(u.Email)).Limit(1).Exist(ctx)
}

func (r *repositroy) CreateUser(ctx context.Context, u *User) (*ent.User, error) {
	return r.client.Create().SetEmail(u.Email).SetPassword(u.Password).SetFirstName(u.FirstName).SetLastName(u.LastName).Save(ctx)
}

func (r *repositroy) GetUser(ctx context.Context, userID uuid.UUID) (*ent.User, error) {
	users, err := r.client.Query().Where(user.UserIDEQ(userID)).Limit(1).All(ctx)
	if err != nil {
		return nil, common.InternalServerError(err)
	}

	if len(users) == 0 {
		return nil, nil
	}

	return users[0], nil
}

func (r *repositroy) UpdateUser(ctx context.Context, userID uuid.UUID, u *User) (int, error) {
	entBuilder := r.client.Update().Where(user.UserIDEQ(userID))

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

	return entBuilder.Save(ctx)
}

func (r *repositroy) DeleteUser(ctx context.Context, userID uuid.UUID) (int, error) {
	return r.client.Delete().Where(user.UserIDEQ(userID)).Exec(ctx)
}
