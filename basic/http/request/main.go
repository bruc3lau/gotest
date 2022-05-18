package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	req, _ := http.NewRequest("GET", "http://baidu.com", nil)
	res, _ := (&http.Client{}).Do(req)
	if b, err := ioutil.ReadAll(res.Body); err == nil {
		fmt.Println(string(b))
	}
}
