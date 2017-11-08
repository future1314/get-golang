package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice("a", a)
	b := make([]int, 0, 5)
	printSlice("b", b)
	c := b[:2]
	printSlice("c", c)
	d := b[1:3]
	printSlice("d", d)
	d = append(d, 1)
	printSlice("d", d)
	d = append(d, 1, 2, 3)
	printSlice("d", d)
}

func printSlice(s string, a []int) {
	fmt.Printf("%s, len=%d, cap=%d\n", s, len(a), cap(a))
}
