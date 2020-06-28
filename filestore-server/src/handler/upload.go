package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/labstack/gommon/log"
	"github.com/leozhao0709/golang/filestore-server/src/db"
	"github.com/leozhao0709/golang/filestore-server/src/handler/handlererror"
	"github.com/leozhao0709/golang/filestore-server/src/hooks"
	"github.com/leozhao0709/golang/filestore-server/src/meta"
	"github.com/leozhao0709/golang/filestore-server/src/util"
)

// UploadHandler handling file uploads
func UploadHandler(w http.ResponseWriter, r *http.Request) *handlererror.HandleError {

	hooks.UseAuth(w, r)

	if r.Method == http.MethodGet {
		bytes, err := ioutil.ReadFile("src/static/view/index.html")
		if err != nil {
			return handlererror.InternalServerError(err)
		}
		w.Write(bytes)
	}
	if r.Method == http.MethodPost {
		file, header, err := r.FormFile("file")
		username := r.FormValue("username")
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

		// meta.UpdateFileMeta(fileMeta)
		err = meta.UpdateFileMetaDB(fileMeta)
		if err != nil {
			return handlererror.InternalServerError(err)
		}

		// update userfile db
		err = db.SaveUserFile(username, fileMeta.FileSha1, fileMeta.FileName, fileMeta.FileSize)
		if err != nil {
			return handlererror.InternalServerError(err)
		}

		http.Redirect(w, r, "/static/view/home.html", http.StatusFound)
	}

	return nil
}

// UploadSuccessHandler Upload successfully hanlder
func UploadSuccessHandler(w http.ResponseWriter, r *http.Request) *handlererror.HandleError {
	w.Write([]byte("Upload successfully!"))
	return nil
}

// FastUploadHandler Fast Upload Handler
func FastUploadHandler(w http.ResponseWriter, r *http.Request) *handlererror.HandleError {

	authErr := hooks.UseAuth(w, r)
	if authErr != nil {
		return authErr
	}

	filehash := r.FormValue("filehash")
	username := r.FormValue("username")
	filename := r.FormValue("filename")

	fileMeta, err := meta.GetFileMetaDB(filehash)

	if err != nil {
		return handlererror.InternalServerError(err)
	}

	if fileMeta == nil {
		json.NewEncoder(w).Encode(util.RespMsg{
			Code: -1,
			Msg:  "Fast Upload file not exist, please use upload api",
		})
		return nil
	}

	err = db.SaveUserFile(username, filehash, filename, fileMeta.FileSize)

	if err != nil {
		json.NewEncoder(w).Encode(util.RespMsg{
			Code: -2,
			Msg:  "Fast Upload fail, please try a few seconds later",
		})
		return nil
	}

	json.NewEncoder(w).Encode(util.RespMsg{
		Code: 0,
		Msg:  "Success",
	})

	return nil
}

// GetFileMetaHandler Get file meta data
func GetFileMetaHandler(w http.ResponseWriter, r *http.Request) *handlererror.HandleError {
	hooks.UseAuth(w, r)
	filehash := r.FormValue("filehash")
	// filemeta, err := meta.GetFileMeta(filehash)
	filemeta, err := meta.GetFileMetaDB(filehash)
	if err != nil {
		// return handlererror.NotFoundError(err)
		log.Error(err)
	}
	data, err := json.Marshal(&filemeta)
	if err != nil {
		return handlererror.InternalServerError(err)
	}
	w.Write(data)

	return nil
}

// QueryUserFileMetasHandler Query all user files with limit and username
func QueryUserFileMetasHandler(w http.ResponseWriter, r *http.Request) *handlererror.HandleError {
	hooks.UseAuth(w, r)
	limit, err := strconv.Atoi(r.FormValue("limit"))
	username := r.FormValue("username")
	if err != nil {
		return handlererror.InternalServerError(err)
	}

	userFileMetas, err := db.QueryUserFileMetas(username, limit)
	if err != nil {
		return handlererror.InternalServerError(err)
	}

	err = json.NewEncoder(w).Encode(userFileMetas)
	if err != nil {
		return handlererror.InternalServerError(err)
	}

	return nil
}

// DownloadHanlder download file handler
func DownloadHanlder(w http.ResponseWriter, r *http.Request) *handlererror.HandleError {
	hooks.UseAuth(w, r)
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
	hooks.UseAuth(w, r)
	if r.Method != http.MethodPost {
		return handlererror.MethodNotAllowedError(r)
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
	hooks.UseAuth(w, r)
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
