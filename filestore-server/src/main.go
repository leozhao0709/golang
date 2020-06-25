package main

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/labstack/gommon/log"
	"github.com/leozhao0709/golang/filestore-server/src/handler"
	"github.com/leozhao0709/golang/filestore-server/src/handler/handlererror"
)

type requestHandler func(http.ResponseWriter, *http.Request)

func handlerWrapper(handlerFunc func(http.ResponseWriter, *http.Request) *handlererror.HandleError) requestHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Debug(r.Method, " ", r.URL.Path)
		err := handlerFunc(w, r)
		if err != nil {
			http.Error(w, http.StatusText(err.StatusCode), err.StatusCode)
			log.Errorf("%v %v %v", r.Method, r.RequestURI, err.Err)
		}
	}
}

func staticServerWrapper(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.RequestURI, "/") {
			http.NotFound(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	log.SetLevel(log.DEBUG)
	log.SetHeader("${time_rfc3339} ${level} ${prefix}")

	// static file
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal("static file failed")
	}
	staticFileServer := http.FileServer(http.Dir(filepath.Join(pwd, "src")))
	http.Handle("/static/", staticServerWrapper(staticFileServer))

	http.HandleFunc("/file/upload", handlerWrapper(handler.UploadHandler))
	http.HandleFunc("/file/upload/success", handlerWrapper(handler.UploadSuccessHandler))
	http.HandleFunc("/file/meta", handlerWrapper(handler.GetFileMetaHandler))
	http.HandleFunc("/file/download", handlerWrapper(handler.DownloadHanlder))
	http.HandleFunc("/file/update", handlerWrapper(handler.FileMetaUpdateHandler))
	http.HandleFunc("/file/delete", handlerWrapper(handler.FileDeleteHandler))

	// user
	http.HandleFunc("/user/signup", handlerWrapper(handler.SignupHandler))
	http.HandleFunc("/user/signin", handlerWrapper(handler.SigninHandler))

	log.Info("server start listening at port 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("sever start fail", err)
	}
}
