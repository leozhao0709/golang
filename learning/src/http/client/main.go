package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {

	req, err := http.NewRequest(http.MethodGet, "http://www.imooc.com", nil)
	req.Header.Add("key", "value")
	resp, err := http.DefaultClient.Do(req)

	// resp, err := http.Get("http://www.imooc.com")
	if err != nil {
		panic(err)
	}

	// must close respons body
	defer resp.Body.Close()

	bytes, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}
