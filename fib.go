package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
     s1 := 0
     s2 := 1
     return func() int {
     	    tmp:= s1
	    s1 = s2
	    s2 = tmp + s1
	    return s1
     }
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}