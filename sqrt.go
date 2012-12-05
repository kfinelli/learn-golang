package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	var z float64
	zn := float64(1)
	for math.Abs(zn-z)>1e-1 {
		z=zn
		zn=z-(z*z-x)/(2*z)
	}
	return zn
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}

