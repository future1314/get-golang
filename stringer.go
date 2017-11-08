package main

import "fmt"

type Person struct {
	name string
	age int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years old)", p.name, p.age)
}

func main() {
	p1 := Person{"zhangsan", 12}
	p2 := Person{"lisi", 16}
	fmt.Println(p1, p2)
}
