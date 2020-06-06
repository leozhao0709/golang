package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", func(writer http.ResponseWriter, req *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(struct {
			Message string `json:"message"`
		}{Message: "pong"})
	})
	http.ListenAndServe(":8080", nil)
}
