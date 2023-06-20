/*
использования select для обработки ошибок при чтении
из канала
*/
package main

import "fmt"

func main() {
	ch := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()

	go func() {
		for {
			select {
			case val, ok := <-ch:
				if !ok {
					done <- true
					return
				}
				fmt.Println("Received:", val)
			}
		}
	}()

	<-done
}
