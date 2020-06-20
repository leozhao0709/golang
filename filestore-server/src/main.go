package main

import (
	"net/http"

	"github.com/leozhao0709/golang/filestore-server/src/handler"
)

func main() {
	http.HandleFunc("/file/upload", handler.UploadHandler)
	http.ListenAndServe(":8080", nil)
}
