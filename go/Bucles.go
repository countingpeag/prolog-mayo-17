package main

import (
	"fmt"
	"math"
)

const delta = 1e-6

func Sqrt(x float64) float64 {
	z := x
	n := 0.0
	for math.Abs(n-z) > delta {
		n, z = z, z-(z*z-x)/(2*z)
	}
	return z
}

func main() {
	const x = 2
	mine, theirs := Sqrt(x), math.Sqrt(x)
	fmt.Println(mine, theirs, mine-theirs)
}
