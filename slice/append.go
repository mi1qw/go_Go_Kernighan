package main

import (
	"fmt"
	"time"
)

func main() {
	var sl []int
	go func() {
		for i := 0; i < 20; i++ {
			fmt.Printf("         %v\n", sl)
			time.Sleep(time.Millisecond * 10)
		}
	}()
	for i := 0; i < 10; i++ {
		sl = append(sl, i)
		fmt.Printf("%-5d %2d   %p\n", len(sl), cap(sl), sl)
		time.Sleep(time.Millisecond * 20)
	}

	/*
	      адрес первого элемента слайса, так как sl само по себе является
	   	указателем на первый элемент внутреннего массива слайса.
	*/

	time.Sleep(time.Second)
	println()
	sl1 := make([]int, 0, 10)
	for i := 0; i < 11; i++ {
		sl1 = append(sl1, i)
		fmt.Printf("%-5d %2d   %p\n", len(sl1), cap(sl1), sl1)
		time.Sleep(time.Millisecond * 20)
	}
}
