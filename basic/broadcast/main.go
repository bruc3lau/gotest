package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println(66)
	var cond sync.Cond
	cond.L.Lock()
	fmt.Println(cond)
}
