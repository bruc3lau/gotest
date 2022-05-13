package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var once sync.Once
	total := 100

	for i := 0; i < total; i++ {
		once.Do(func() {
			fmt.Println(time.Now())
		})
	}
}
