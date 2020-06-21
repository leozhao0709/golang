package meta

import "fmt"

// FileMeta File metadata
type FileMeta struct {
	FileSha1 string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}

var fileMetas map[string]FileMeta

func init() {
	fileMetas = make(map[string]FileMeta)
}

// UpdateFileMeta add or update file meta data
func UpdateFileMeta(fmeta FileMeta) {
	fileMetas[fmeta.FileSha1] = fmeta
}

// GetFileMeta get the file meta data with sha
func GetFileMeta(fileSha1 string) (FileMeta, error) {
	meta, ok := fileMetas[fileSha1]
	if !ok {
		return meta, fmt.Errorf("fileSha1 %v cannot be found", fileSha1)
	}
	return meta, nil
}

// RemoveFileMeta remove one file meta with sha1
func RemoveFileMeta(fileSha1 string) (FileMeta, error) {
	filemeta, err := GetFileMeta(fileSha1)

	if err != nil {
		return filemeta, err
	}

	delete(fileMetas, fileSha1)
	return filemeta, nil
}
