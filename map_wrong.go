package main

import (
	"code.google.com/p/go-tour/wc"
	"fmt"
	"strings"
)

func CountToken(s,j string) int {
     count:=0
     	 index:=0
	 for {
	     fmt.Println(index)

	     tmpindex:=strings.Index(s[index:],j)
	     if tmpindex==-1 {
	     return count
	     } else {
	     index += tmpindex
	     }

	     if index==0 && s[index+len(j)]==' ' {
	     	  count++
		  index+=len(j)
	     } else if index>0 && s[index-1]==' ' && (index+len(j)==len(s) || s[index+len(j)]==' ') {
	     	  count++
		  index+=len(j)
	     } else {
	     index += len(j)
	     }
	 }
	 return count
}

func WordCount(s string) map[string]int {
     outmap := make(map[string]int)
     //so Fields doesn't do what I thought it did... 
     //if it returned a non-repeating list (set) of 
     //tokens, then I think this would be the way to 
     //solve this
     for _,j := range strings.Fields(s) {
     	 fmt.Println(j)
	 outmap[j]+=CountToken(s,j)
     }
     return outmap
}



func main() {
	wc.Test(WordCount)
	fmt.Println(" ")
}