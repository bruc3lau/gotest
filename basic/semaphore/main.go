package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func get(id string) string {
	response, _ := http.Get(fmt.Sprintf("https://jsonplaceholder.typicode.com/todos/%s", id))
	res, _ := io.ReadAll(response.Body)
	// content := &(string(res))
	// fmt.Println(*content)
	fmt.Println(string(res))
	return string(res)
}

func main() {
	for i := 1; i < 2; i++ {
		go get(strconv.Itoa(i))
	}
}
