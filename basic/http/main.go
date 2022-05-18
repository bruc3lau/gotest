package main

import (
	_ "bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// req, _ := http.NewRequest("GET", "baidu.com", new(bytes.Buffer))
	// res, _ := (&http.Client{}).Do(req)
	// fmt.Println(res.Body)
	res, err := http.Get("http://baidu.com")
	fmt.Printf("%+v,%s\n", res.Body, err)
	bs, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(bs))
}
