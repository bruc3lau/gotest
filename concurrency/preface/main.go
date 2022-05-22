package main

import (
	"fmt"
	"sync"
	"time"
	_ "time"
)

func main() {
	// var data int
	// go func() {
	// 	data++
	// }()
	// time.Sleep(time.Microsecond)
	// if data == 0 {
	// 	fmt.Printf("the value is %v.\n", data)
	// }

	mux()
}

func mux() {
	var memoryAccess sync.Mutex
	var value int
	go func() {
		memoryAccess.Lock()
		value++
		memoryAccess.Unlock()
	}()
	time.Sleep(time.Microsecond)
	memoryAccess.Lock()
	if value == 0 {
		fmt.Printf("the value is %v.\n", value)
	} else {
		fmt.Printf("the value is %v.\n", value)
	}
	memoryAccess.Unlock()
}
