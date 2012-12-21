package main

import "fmt"
import "math/cmplx"

func Cbrt(x complex128) complex128 {
	var z complex128

	//cannot use 1+0i as a starting value, 1+i works 
	//for a few tested values
	//
	//should think about which starting values will 
	//converge for different radicands

	zn := complex128(1+1i) 
	for cmplx.Abs(zn-z)>1e-3 {
		z=zn
		zn=z-(z*z*z-x)/(3*z*z)
	}
	return zn

}

func main() {
	fmt.Println(Cbrt(-2))
	fmt.Println(cmplx.Pow(-2,1./3.))
}