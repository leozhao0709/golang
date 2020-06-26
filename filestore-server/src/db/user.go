package db

import (
	"time"

	"github.com/labstack/gommon/log"
)

// User user model
type User struct {
	Username     string    `db:"user_name"`
	Email        string    `db:"email"`
	Phone        string    `db:"phone"`
	SignupAt     time.Time `db:"signup_at"`
	LastActiveAt time.Time `db:"last_active"`
	Status       int       `db:"status"`
}

// UserSignup user sign up
func UserSignup(username string, password string) error {
	result, err := GetDB().Exec("insert into tbl_user (`user_name`, `user_pwd`) values(?, ?)", username, password)

	if err != nil {
		log.Info("user sign up fail ", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Info("query rowsAffected error")
	}

	log.Info("user sign up with rowsAffected: ", rowsAffected)

	return nil
}

// UserSignin user sign in check
func UserSignin(username string, shaPassword string) (bool, error) {

	var dbShaPassword string
	err := GetDB().Get(&dbShaPassword, "select user_pwd from tbl_user where user_name=? limit 1", username)

	if err != nil {
		return false, err
	}

	if shaPassword != dbShaPassword {
		return false, nil
	}

	return true, nil
}

// UpdateUserToken update user token
func UpdateUserToken(username string, token string) error {
	_, err := GetDB().Exec("insert into tbl_user_token (user_name, user_token) value(?, ?) on duplicate key update user_token=?", username, token, token)

	if err != nil {
		return err
	}
	return nil
}

// GetUserInfo get user info
func GetUserInfo(username string) (User, error) {
	var user User
	err := GetDB().Get(&user, "select user_name, email, phone, signup_at, last_active, status from tbl_user where user_name=?", username)
	return user, err
}
