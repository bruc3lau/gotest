package main

import (
	"fmt"
	"sync"
	"time"
)

func doWork(i int) {
	time.Sleep(time.Second)
	fmt.Println("current:", i)
}

func main() {
	var wg sync.WaitGroup
	total := 100

	// wg.Add(total)

	// fmt.Println(1)
	for i := 0; i < total; i++ {
		wg.Add(1)
		// fmt.Println(i)
		go func(i int) {
			time.Sleep(time.Second)
			fmt.Println("current:", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
