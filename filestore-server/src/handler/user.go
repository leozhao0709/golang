package handler

import (
	"io/ioutil"
	"net/http"

	"github.com/leozhao0709/golang/filestore-server/src/handler/handlererror"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) *handlererror.HandleError {
	if r.Method == http.MethodGet {
		bytes, err := ioutil.ReadFile("./src/static/view/signup.html")
		if err != nil {
			return handlererror.InternalServerError(err)
		}
		w.Write(bytes)
	}
	return nil
}
