package main

import (
	"fmt"
	"time"
)

func doWork() {
	time.Sleep(time.Second)
	fmt.Println("done!")
}

func main() {
	for i := 1; i < 10; i++ {
		defer doWork()
	}
}
