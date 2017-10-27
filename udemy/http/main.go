package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://google.com")
	if err != nil {
		log.Fatalln("Error", err)
	}

	// bs := make([]byte, 99999)
	// lenth, _ := res.Body.Read(bs)
	// log.Println(string(bs))
	// log.Println(lenth)

	// io.Copy(os.Stdout, res.Body)

	lw := logWriter{}
	io.Copy(lw, res.Body)
}

type logWriter struct{}

func (logWriter) Write(bs []byte) (int, error) {
	log.Println(string(bs))
	log.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
