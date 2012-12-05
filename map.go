package main

import (
	"code.google.com/p/go-tour/wc"
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
     outmap := make(map[string]int)
     for _,j := range strings.Fields(s) {
	 outmap[j]++
     }
     return outmap
}



func main() {
	wc.Test(WordCount)
	fmt.Println(" ")
}