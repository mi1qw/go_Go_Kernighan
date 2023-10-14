package main

import (
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Wait()

	wg.Add(1)
	wg.Done()
	wg.Wait()
	println("Hello")
}
