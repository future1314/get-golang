package main

import (
	"math"
	"fmt"
)

type Abser interface {
	Abs() float64
}

func main() {
	f := MyFloat1(-2.3)
	v := Vertex1{3, 4}

	var a Abser

	a = f
	fmt.Println(a.Abs())

	a = &v
	fmt.Println(a.Abs())
}

type MyFloat1 float64

func (f MyFloat1) Abs() float64 {
	if f < 0 {
		return -float64(f)
	}
	return float64(f)
}

type Vertex1 struct {
	x, y float64
}

func (v *Vertex1) Abs() float64 {
	return math.Sqrt(v.x * v.x + v.y * v.y)
}
