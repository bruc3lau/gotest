package main

import (
	"fmt"
	"time"
)

func doWork() {
	time.Sleep(time.Second)
	fmt.Println(time.Now())
}

func main() {
	for i := 0; i < 10000; i++ {
		go doWork()
	}
	select {
	case <-time.After(time.Second):

	}
}
