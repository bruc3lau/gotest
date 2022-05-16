package main

import (
	"context"
	"fmt"
	"time"
)

func doWork() {
	time.Sleep(time.Second * 2)
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	fmt.Println(ctx)
	fmt.Println(cancel)
	fmt.Println(time.Now())
	fmt.Println(time.Now(), <-ctx.Done())
	doWork()
	fmt.Println(time.Now(), <-ctx.Done())
	fmt.Println(time.Now(), <-ctx.Done())
}
