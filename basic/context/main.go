package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	type myKey string
	ctx := context.Background()
	fmt.Println(ctx)

	//value
	ctx2 := context.WithValue(ctx, "a", 1)
	ctx3 := context.WithValue(ctx2, myKey("hello"), "world")
	fmt.Println(ctx2)
	fmt.Println(ctx3)

	fmt.Println(ctx3.Value("a"))
	fmt.Println(ctx3.Value(myKey("hello")))

	fmt.Printf("%+v\n", &ctx2)
	fmt.Printf("%v\n", &ctx2)
	fmt.Printf("%v\n", &ctx3)

	ctx4, cancel := context.WithTimeout(ctx3, time.Millisecond*100)
	defer cancel()
	fmt.Println(ctx4)

	time.Sleep(time.Millisecond * 1000)

}
