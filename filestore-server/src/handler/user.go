package handler

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/leozhao0709/golang/filestore-server/src/db"
	"github.com/leozhao0709/golang/filestore-server/src/handler/handlererror"
	"github.com/leozhao0709/golang/filestore-server/src/util"
)

const userPasswordSalt = "12345"

// SignupHandler user sign up handler
func SignupHandler(w http.ResponseWriter, r *http.Request) *handlererror.HandleError {
	if r.Method == http.MethodGet {
		bytes, err := ioutil.ReadFile("./src/static/view/signup.html")
		if err != nil {
			return handlererror.InternalServerError(err)
		}
		w.Write(bytes)
	}

	if r.Method == http.MethodPost {
		username := strings.TrimSpace(r.PostFormValue("username"))
		password := strings.TrimSpace(r.PostFormValue("password"))

		if len(username) < 3 || len(password) < 5 {
			return handlererror.BadRequestError(errors.New("username or password is not match with requirement"))
		}

		shaPassword := util.Sha1([]byte(password + userPasswordSalt))
		err := db.UserSignup(username, shaPassword)
		if err != nil {
			return handlererror.InternalServerError(err)
		}

		w.Write([]byte("SUCCESS"))
	}
	return nil
}
