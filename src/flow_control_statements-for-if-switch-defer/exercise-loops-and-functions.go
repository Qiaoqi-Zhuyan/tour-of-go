package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	for i := 0; i < 100; i++ {
		z -= (z*z - x) / (x * z)
		fmt.Printf("z is : %v\n", z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(9))
}
