package hooks

import (
	"errors"
	"net/http"

	"github.com/leozhao0709/golang/filestore-server/src/handler/handlererror"
	"github.com/leozhao0709/golang/filestore-server/src/util"
)

// UseAuth user Auth Hook
func UseAuth(w http.ResponseWriter, r *http.Request) *handlererror.HandleError {
	username := r.FormValue("username")
	token := r.FormValue("token")

	if len(username) < 3 || !util.IsTokenValid(token) {
		// http.Redirect(w, r, "/user/signin", http.StatusFound)
		return handlererror.UnauthorizedRequestError(errors.New("username or usertoken is invalid"))
	}

	return nil
}
