package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	res, prev, curr := 0, 0, 1
	fib := func() int {
		res, prev, curr = prev, curr, prev+curr
		return res
	}

	return fib
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
