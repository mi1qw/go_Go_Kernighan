package main

import "fmt"

func main() {
	var sl []int
	for i := 0; i < 10; i++ {
		sl = append(sl, i)
		//println(&sl, cap(sl))
		fmt.Printf("%-5d %d\n", len(sl), cap(sl))
	}
}
