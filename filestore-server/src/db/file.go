package db

import (
	"database/sql"
	"time"

	"github.com/labstack/gommon/log"
)

// SaveFile Save File to DB
func SaveFile(filehash string, filename string, filesize int64, fileaddr string) error {
	result, err := GetDB().NamedExec("insert into tbl_file (file_sha1, file_name, file_size, file_addr, status) values(:filehash, :filename, :filesize, :fileaddr, :status)", map[string]interface{}{
		"filehash": filehash,
		"filename": filename,
		"filesize": filesize,
		"fileaddr": fileaddr,
		"status":   1,
	})

	if err != nil {
		return err
	}

	count, err := result.RowsAffected()
	log.Debug("SaveFile affected row count is ", count)
	if count <= 0 {
		log.Infof("File with hash %v is same with before", filehash)
	}

	return err
}

// TableFile file table
type TableFile struct {
	FileHash string        `db:"file_sha1"`
	FileName string        `db:"file_name"`
	FileSize sql.NullInt64 `db:"file_size"`
	FileAddr string        `db:"file_addr"`
	UploadAt *time.Time    `db:"update_at"`
}

// GetFileMeta get the file meta from db
func GetFileMeta(filehash string) (*TableFile, error) {
	var tableFile = &TableFile{}
	err := GetDB().Get(tableFile, "select file_sha1, file_name, file_size, file_addr, update_at from tbl_file where file_sha1=? and status=1", filehash)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return tableFile, nil
}
