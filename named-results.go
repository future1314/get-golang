package main

import "fmt"

func swap(a int) (x, y int) {
	x = a + 1
	y = a - 1
	return
}

func main() {
	x, y := swap(5)
	fmt.Println(x, y)
}
