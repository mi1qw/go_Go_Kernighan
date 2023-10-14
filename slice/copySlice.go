package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4}
	news := s[1:3]
	fmt.Printf("%v \n", news) // [2 3]

	ints := make([]int, 2, 5)
	fmt.Printf("len %v  cap %v \n\n", len(ints), cap(ints))

	ints[1] = 1
	ints = append(ints, 2)
	fmt.Printf("len %v  cap %v \n%v\n\n", len(ints), cap(ints), ints)

	var a [3]int
	fmt.Printf("len %v  cap %v \n%v\n\n", len(a), cap(a), a)
}
