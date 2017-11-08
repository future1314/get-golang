package main

import (
	"fmt"
	"time"
)

func Sum(x []int, c chan int) {
	sum := 0
	for _, v := range x {
		time.Sleep(time.Second)
		fmt.Printf("%v ", v)
		sum += v
	}
	c <- sum
}

func main() {
	c := make(chan int)
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	go Sum(a[:len(a)/3], c)
	go Sum(a[len(a)/3:], c)
	x, y := <-c, <-c

	fmt.Println()
	fmt.Println(x, y, x + y)
}
