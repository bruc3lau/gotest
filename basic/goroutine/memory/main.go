package main

import "time"

func doWork() {
	time.Sleep(time.Minute * 60)
}

func main() {
	for i := 0; i < 100; i++ {
		go doWork()
	}
	time.Sleep(time.Minute * 60)
}
