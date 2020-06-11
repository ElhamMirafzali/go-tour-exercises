package main

import (
	"fmt"
	"math"
)

func Sqrt1(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func Sqrt2(x float64) float64 {
	z := 1.0
	for math.Abs((z*z)-x) > 1e-15 {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Printf("Math.Sqrt(%v):%v\n", 2, math.Sqrt(2))
	fmt.Printf("Sqrt1(%v):%v\n", 2, Sqrt1(2))
	fmt.Printf("Sqrt2(%v):%v", 2, Sqrt2(2))
}
