package main

import "fmt"

type Point struct {
	x int
	y int
}

var (
	a = Point{1, 2}
	b = Point{x: 3}
	c = Point{}
	p1 = &Point{6, 7}
)

func main() {
	v := Point{3, 5}
	p := &v
	p.x = 1e9
	fmt.Println(v, a, b, c, p1)
}
