package main

import (
	"time"
	"fmt"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(1000 * time.Millisecond)

	for {
		select {
		case <- tick:
			fmt.Println("tick")
		case <- boom:
			fmt.Println("boom!")
			return
		default:
			fmt.Println("   .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
