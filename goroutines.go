package main

import (
	"time"
	"fmt"
)

func SayHello(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go SayHello("world")
	SayHello("Hello")
}
