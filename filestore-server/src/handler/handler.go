package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/labstack/gommon/log"
	"github.com/leozhao0709/golang/filestore-server/src/handler/handlererror"
	"github.com/leozhao0709/golang/filestore-server/src/meta"
	"github.com/leozhao0709/golang/filestore-server/src/util"
)

// UploadHandler handling file uploads
func UploadHandler(w http.ResponseWriter, r *http.Request) *handlererror.HandleError {

	if r.Method == http.MethodGet {
		bytes, err := ioutil.ReadFile("src/static/view/index.html")
		if err != nil {
			return handlererror.InternalServerError(err)
		}
		w.Write(bytes)
	}
	if r.Method == http.MethodPost {
		file, header, err := r.FormFile("file")
		if err != nil {
			return handlererror.InternalServerError(err)
		}
		defer file.Close()

		fileMeta := meta.FileMeta{
			FileName: header.Filename,
			Location: "/tmp/" + header.Filename,
			UploadAt: time.Now().Format("2006-01-02 15:04:05"),
		}

		newFile, err := os.Create(fileMeta.Location)
		if err != nil {
			return handlererror.InternalServerError(err)
		}

		fileSize, err := io.Copy(newFile, file)
		if err != nil {
			return handlererror.InternalServerError(err)
		}

		fileMeta.FileSize = fileSize

		// must seek before you get filesha1
		newFile.Seek(0, 0)
		fileMeta.FileSha1 = util.FileSha1(newFile)
		log.Debugf("stored filemeta is %+v", fileMeta)

		// meta.UpdateFileMeta(fileMeta)
		err = meta.UpdateFileMetaDB(fileMeta)
		if err != nil {
			return handlererror.InternalServerError(err)
		}

		http.Redirect(w, r, "/file/upload/success", http.StatusFound)
	}

	return nil
}

// UploadSuccessHandler Upload successfully hanlder
func UploadSuccessHandler(w http.ResponseWriter, r *http.Request) *handlererror.HandleError {
	w.Write([]byte("Upload successfully!"))
	return nil
}

// GetFileMetaHandler Get file meta data
func GetFileMetaHandler(w http.ResponseWriter, r *http.Request) *handlererror.HandleError {
	filehash := r.FormValue("filehash")
	log.Debug("file hash is ", filehash)
	filemeta, err := meta.GetFileMeta(filehash)
	if err != nil {
		return handlererror.NotFoundError(err)
	}
	data, err := json.Marshal(&filemeta)
	if err != nil {
		return handlererror.InternalServerError(err)
	}
	w.Write(data)

	return nil
}

// DownloadHanlder download file handler
func DownloadHanlder(w http.ResponseWriter, r *http.Request) *handlererror.HandleError {
	fileSha1 := r.FormValue("filehash")
	log.Debug("file hash is ", fileSha1)
	filemeta, err := meta.GetFileMeta(fileSha1)
	if err != nil {
		return handlererror.NotFoundError(err)
	}

	file, err := os.Open(filemeta.Location)
	if err != nil {
		return handlererror.InternalServerError(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return handlererror.InternalServerError(err)
	}
	w.Header().Set("Content-Type", "application/octect-stream")
	w.Header().Set("content-disposition", fmt.Sprintf(`attachment;filename="%s"`, filemeta.FileName))
	w.Write(data)
	return nil
}

// FileMetaUpdateHandler update file name
func FileMetaUpdateHandler(w http.ResponseWriter, r *http.Request) *handlererror.HandleError {

	if r.Method != http.MethodPost {
		return handlererror.MethodNotAllowedError(fmt.Errorf("%v is only allowed Post", r.RequestURI))
	}

	opType := r.FormValue("opType")
	fileSha1 := r.FormValue("filehash")
	newFileName := r.FormValue("filename")

	if opType != "0" {
		return handlererror.ForbiddenError(errors.New("op cannot be a non 0 value"))
	}

	filemeta, err := meta.GetFileMeta(fileSha1)
	if err != nil {
		return handlererror.NotFoundError(err)
	}

	filemeta.FileName = newFileName
	meta.UpdateFileMeta(filemeta)

	data, err := json.Marshal(&filemeta)
	if err != nil {
		return handlererror.InternalServerError(err)
	}

	w.Write(data)

	return nil
}

// FileDeleteHandler delete file handler
func FileDeleteHandler(w http.ResponseWriter, r *http.Request) *handlererror.HandleError {
	fileSha1 := r.FormValue("filehash")

	filemeta, err := meta.RemoveFileMeta(fileSha1)
	if err != nil {
		return handlererror.NotFoundError(err)
	}

	err = os.Remove(filemeta.Location)
	if err != nil {
		return handlererror.InternalServerError(err)
	}

	return nil
}
