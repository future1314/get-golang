package main

import "fmt"

func main() {
	sum1 := 0
	for i:=0; i<10; i++ {
		sum1 += i
	}
	sum2 := 1
	for ; sum2 < 100; {
		sum2 += sum2
	}
	sum3 := 1
	for sum3 < 100 {
		sum3 += sum3
	}
	fmt.Println(sum1, sum2, sum3)
}
