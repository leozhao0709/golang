package db

import (
	"database/sql"

	"github.com/leozhao0709/golang/filestore-server/src/common/formattime"
)

// UserFileMeta User File from db
type UserFileMeta struct {
	ID         int
	Username   string          `db:"user_name"`
	FileHash   string          `db:"file_sha1"`
	FileSize   int64           `db:"file_size"`
	FileName   string          `db:"file_name"`
	UploadAt   formattime.Time `db:"upload_at"`
	LastUpdate formattime.Time `db:"last_update"`
	Status     int
}

// SaveUserFile Save user file
func SaveUserFile(username, filehash, filename string, filesize int64) error {

	_, err := GetDB().Exec("insert into tbl_user_file (user_name, file_sha1, file_size, file_name) values(?, ?, ?, ?)", username, filehash, filesize, filename)

	if err != nil {
		return err
	}

	return nil
}

// QueryUserFileMetas Query user file metas with username and limit
func QueryUserFileMetas(username string, limit int) ([]UserFileMeta, error) {
	userFileMetas := []UserFileMeta{}

	// err := GetDB().Select(&userFileMetas, "select user_name, file_sha1, file_size, file_name, upload_at, last_update, status from tbl_user_file where user_name=? limit ?", username, limit)

	err := GetDB().Select(&userFileMetas, "select * from tbl_user_file where user_name=? limit ?", username, limit)

	if err != nil && err != sql.ErrNoRows {
		return userFileMetas, err
	}

	return userFileMetas, nil
}
