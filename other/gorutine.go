package main

import "time"

func main() {
	go func() {
		println("start 1")

		go func() {
			println("start 2")
			time.Sleep(time.Second)
			println("end 2")
		}()

		println("end 1")
	}()
	time.Sleep(time.Second * 2)
}
