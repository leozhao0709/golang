package handler

import (
	"io/ioutil"
	"net/http"
)

// UploadHandler handling file uploads
func UploadHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		bytes, err := ioutil.ReadFile("src/static/view/upload.html")
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		writer.Write(bytes)
	}
}
