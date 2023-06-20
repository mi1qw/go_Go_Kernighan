/*
В этом примере программа создает канал ch, на который горутина отправляет значения
в течение 2.5 секунд, а затем посылает сигнал о завершении работы в канал done.
В главной функции программа использует оператор select, чтобы либо принимать значения
из канала ch, либо прерваться, если в канал done поступит сигнал.
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			time.Sleep(500 * time.Millisecond)
		}
		done <- true
	}()

	for {
		select {
		case val := <-ch:
			fmt.Println("Received:", val)
		case <-done:
			fmt.Println("Done!")
			return
		}
	}
}
