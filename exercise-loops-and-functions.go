package main

import(
	"fmt"
	"math"
)

func newton_sqrt(x float64) float64 {
	z := 1.0
//	for i := 0; i < 10; i++ {
//		z = z - ((z * z) - x) / (2 * z)
//	}
//	return z

	for {
		if y := ((z * z) - x) / (2 * z); math.Abs(y) > 10e-10 {
			z = z - y
			fmt.Println(y, z)
		} else {
			return z
		}
	}
}

func main() {
	fmt.Println(newton_sqrt(2), math.Sqrt(2))
}
