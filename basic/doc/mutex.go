package main

import (
	"fmt"
	"sync"
	"time"
)

func add() {

}

func main() {
	var mu sync.Mutex
	n := 0
	for i := 0; i < 1000; i++ {
		go func(n *int) {
			fmt.Println(mu.TryLock())
			// mu.Lock()
			*n++
			// mu.Unlock()
			// fmt.Println(*n)
		}(&n)
	}
	time.Sleep(time.Second)
	fmt.Println("n: ", n)
}
