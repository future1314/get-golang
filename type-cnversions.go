package main

import(
	"fmt"
	"math"
)

func main() {
	var a, b int = 3, 4
	t := math.Sqrt(float64(a*a+b*b))
	c := int(t)
	fmt.Println(a, b, c)
}
