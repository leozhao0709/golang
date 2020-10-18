package users

import (
	"time"

	"github.com/google/uuid"
)

// User ...
type User struct {
	ID         int       `json:"id"`
	UserID     uuid.UUID `json:"user_id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Email      string    `json:"email"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
	Status     string    `json:"status"`
	Password   string    `json:"password"`
}

// PublicUser ...
type PublicUser struct {
	UserID     uuid.UUID `json:"user_id"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
	Status     string    `json:"status"`
}

// PrivateUser ...
type PrivateUser struct {
	UserID     uuid.UUID `json:"user_id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Email      string    `json:"email"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
	Status     string    `json:"status"`
}
