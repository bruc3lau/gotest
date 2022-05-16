package main

import (
	"fmt"
	"time"
)

func main() {
	select {
	case r := <-time.After(time.Second):
		fmt.Println("hello",r)
		fmt.Printf("%T",r)
		// default:
		// 	fmt.Println("ok")
	}
}
