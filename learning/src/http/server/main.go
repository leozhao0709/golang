package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/ping", func(writer http.ResponseWriter, req *http.Request) {
		<-time.After(5 * time.Second)
		fmt.Println(".......")
		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(struct {
			Message string `json:"message"`
		}{Message: "pong"})
	})
	http.ListenAndServe(":8080", nil)
}
