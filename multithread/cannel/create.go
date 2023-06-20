package main

import (
	"fmt"
	"time"
)

func main() {
	var ch0 chan int
	ch0 = make(chan int)
	print(ch0)

	ch1 := make(chan int, 1)
	print(ch1)

	ch1 <- 5
	print(ch1, "add 5")

	select {
	case ch1 <- 7:
		println(ch1, "add 7")
	case <-time.After(2 * time.Second):
		print(ch1, "blocked")
	}
	<-ch1
	print(ch1, "del 5")
}

func print(ch chan int, str ...string) {
	var txt string
	if len(str) > 0 {
		txt = str[0]
	} else {
		txt = ""
	}
	fmt.Printf("cap=%d  len=%d  %s \n", cap(ch), len(ch), txt)
}
