package main

import (
	"net/http"

	"github.com/labstack/gommon/log"
	"github.com/leozhao0709/golang/filestore-server/src/handler"
	"github.com/leozhao0709/golang/filestore-server/src/handler/handlererror"
)

type requestHandler func(http.ResponseWriter, *http.Request)

func handlerWrapper(handlerFunc func(http.ResponseWriter, *http.Request) *handlererror.HandleError) requestHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Debug(r.Method, " ", r.RequestURI)
		err := handlerFunc(w, r)
		if err != nil {
			http.Error(w, http.StatusText(err.StatusCode), err.StatusCode)
			log.Errorf("%v %v %v", r.Method, r.RequestURI, err.Err)
		}
	}
}

func main() {
	log.SetLevel(log.DEBUG)
	log.SetHeader("${time_rfc3339} ${level} ${prefix}")

	http.HandleFunc("/file/upload", handlerWrapper(handler.UploadHandler))
	http.HandleFunc("/file/upload/success", handlerWrapper(handler.UploadSuccessHandler))
	http.HandleFunc("/file/meta", handlerWrapper(handler.GetFileMetaHandler))
	http.HandleFunc("/file/download", handlerWrapper(handler.DownloadHanlder))
	http.HandleFunc("/file/update", handlerWrapper(handler.FileMetaUpdateHandler))
	http.HandleFunc("/file/delete", handlerWrapper(handler.FileDeleteHandler))

	// user
	http.HandleFunc("/user/signup", handlerWrapper(handler.SignupHandler))

	log.Info("server start listening at port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("sever start fail", err)
	}
}
