package mydb

import "github.com/labstack/gommon/log"

// SaveFile Save File to DB
func SaveFile(filehash string, filename string, filesize int64, fileaddr string) error {
	log.Debug("filehash=", filehash, " filename=", filename, " filesize=", filesize, " fileaddr=", fileaddr, " db=", GetDB())
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
	if count <= 0 {
		log.Infof("File with hash %v is same with before", filehash)
	}

	return err
}
