package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("hello")
	var c http.Client
	fmt.Println(c)
	x, x1 := c.Get("https://api.github.com/repos/qnblackcat/uYouPlus/releases")
	fmt.Println(x, x1)
	b := x.Body
	fmt.Println(b)
	buff := new(bytes.Buffer)
	fmt.Println(buff)
	x2, x3 := buff.ReadFrom(b)
	fmt.Println(x2, x3)
	fmt.Println(buff)

	var result []map[string]interface{}
	fmt.Println(result)
	x4 := json.Unmarshal(buff.Bytes(), &result)
	fmt.Println(x4)
	// fmt.Println(result)

	latest := result[0]
	fmt.Println(latest)
	fmt.Println()
	created_at := latest["created_at"]
	fmt.Println(created_at)
}
