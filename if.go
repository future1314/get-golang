package main

import (
	"fmt"
	"math"
)

func sqrt1(i float64) string {
	if i < 0 {
		return sqrt1(-i) + "i"
	}
	return fmt.Sprint(math.Sqrt(i))
}

func main() {
	fmt.Println(sqrt1(2), sqrt1(-4))
}
