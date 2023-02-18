// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 135.

// The squares program demonstrates a function value with state.
package main

import "fmt"

// !+
// squares returns a function that returns
// the next square number each time it is called.
func squares1() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func squares(val int) func() int {
	x := val
	return func() int {
		x++
		return x * x
	}
}
func main() {
	f := squares1()
	fmt.Println(f()) // "1"
	fmt.Println(f()) // "4"
	fmt.Println(f()) // "9"
	fmt.Println(f()) // "16"
	fmt.Printf("\n\n")

	// list of structures with functions
	var sl []square
	for _, v := range [...]int{1, 3, 7} {
		sl = append(sl, square{v, squares(v)})
	}
	for _, fSq := range sl {
		fmt.Printf("step %d  %d %d\n",
			fSq.i, fSq.f(), fSq.f())
	}
}

// !-
type square struct {
	i int
	f func() int
}
