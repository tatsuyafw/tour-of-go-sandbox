package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	count := 10

	for i := 0; i < count; i++ {
		z = z - ((z*z)-x)/(2*z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(3))
}
