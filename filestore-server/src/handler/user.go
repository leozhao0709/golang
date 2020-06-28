package handler

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/leozhao0709/golang/filestore-server/src/db"
	"github.com/leozhao0709/golang/filestore-server/src/handler/handlererror"
	"github.com/leozhao0709/golang/filestore-server/src/hooks"
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
		return nil
	} else if r.Method == http.MethodPost {
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
		return nil
	} else {
		return handlererror.MethodNotAllowedError(r)
	}
}

// SigninHandler user sign in handler
func SigninHandler(w http.ResponseWriter, r *http.Request) *handlererror.HandleError {

	if r.Method == http.MethodGet {
		bytes, err := ioutil.ReadFile("./src/static/view/signin.html")
		if err != nil {
			return handlererror.InternalServerError(err)
		}
		w.Write(bytes)
		return nil
	} else if r.Method == http.MethodPost {
		username := strings.TrimSpace(r.PostFormValue("username"))
		password := strings.TrimSpace(r.PostFormValue("password"))

		shaPassword := util.Sha1([]byte(password + userPasswordSalt))
		exist, err := db.UserSignin(username, shaPassword)
		if err != nil {
			return handlererror.InternalServerError(err)
		}

		if !exist {
			w.Write([]byte("FAILED"))
			return nil
		}

		// 2. generate token
		token := generateToken(username)
		err = db.UpdateUserToken(username, token)
		if err != nil {
			return handlererror.InternalServerError(err)
		}

		json.NewEncoder(w).Encode(util.RespMsg{
			Code: 0,
			Msg:  "OK",
			Data: struct {
				Location string
				Username string
				Token    string
			}{
				Location: "http://" + r.Host + "/static/view/home.html",
				Username: username,
				Token:    token,
			},
		})

		return nil
	} else {
		return handlererror.MethodNotAllowedError(r)
	}
}

// UserInfoHandler user info handler
func UserInfoHandler(w http.ResponseWriter, r *http.Request) *handlererror.HandleError {
	err := hooks.UseAuth(w, r)
	if err != nil {
		return err
	}

	if r.Method == http.MethodPost {
		username := strings.TrimSpace(r.FormValue("username"))
		user, err := db.GetUserInfo(username)
		if err != nil {
			return handlererror.InternalServerError(err)
		}

		json.NewEncoder(w).Encode(util.RespMsg{
			Code: 0,
			Msg:  "OK",
			Data: user,
		})

		return nil
	}
	return handlererror.MethodNotAllowedError(r)
}

func generateToken(username string) string {
	ts := strconv.Itoa(int(time.Now().Unix()))
	return util.MD5([]byte(username+ts+"_tokensalt")) + ts[:8]
}
