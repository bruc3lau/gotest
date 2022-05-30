package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type resp struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	res := "{'id':1,'name':'Bruce'}"
	b := ``
	fmt.Println(b)
	var r resp
	d := json.NewDecoder(strings.NewReader(res)).Decode(&r)
	fmt.Println(d)
}
