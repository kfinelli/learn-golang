package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var outpic [][]uint8
	for i:=0; i<dy; i++ {
		var yslice []uint8
		for j:=0; j<dx; j++ {
			yslice=append(yslice,uint8(i*j))
		}
		outpic=append(outpic,yslice)
	}
	return outpic
}

func main() {
	pic.Show(Pic)
}