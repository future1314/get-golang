package main

import "fmt"

func main() {
	s := []int{1,2,3,4,5,6,7,8,9}
	fmt.Println(s)

	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d] = %d\n", i, s[i])
	}
}
