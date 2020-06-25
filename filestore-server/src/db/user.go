package db

import "github.com/labstack/gommon/log"

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
