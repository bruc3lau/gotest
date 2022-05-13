package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	ch := make(chan int)

	for i := 0; i < 10; i++ {
		go func() {
			once.Do(func() {
				fmt.Println("print only once!")
			})
			ch <- i
		}()
	}

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}
