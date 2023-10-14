package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func printFirstElement(s []int, cond *sync.Cond, sem chan struct{}) {
	cond.L.Lock()

	for {
		cond.Wait()
		fmt.Printf("%d\n", s[0])
		sem <- struct{}{}
	}

	cond.L.Unlock()
}

func main() {
	cond := sync.NewCond(&sync.RWMutex{})
	s := make([]int, 1)
	sem := make(chan struct{}, runtime.NumCPU())

	for i := 0; i < runtime.NumCPU(); i++ {
		go printFirstElement(s, cond, sem)
	}

	time.Sleep(time.Millisecond * 10)

	for i := 0; i < 5; i++ {
		cond.L.Lock()
		s[0] = i
		cond.Broadcast()
		cond.L.Unlock()

		// Ожидание завершения всех горутин перед переходом к следующему значению i
		for j := 0; j < runtime.NumCPU(); j++ {
			<-sem
		}
	}

	close(sem)
	time.Sleep(time.Millisecond * 300)
}
