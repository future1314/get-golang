package main

import "fmt"

func swap1(a, b string) (string, string) {
	return b, a
}

func main() {
	a, b := swap1("hello", "world")
	fmt.Println(a, b)
}
