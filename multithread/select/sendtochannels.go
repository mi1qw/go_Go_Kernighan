/*
Cоздает два канала ch1 и ch2, а затем запускает горутину,
которая отправляет значения на оба канала по очереди, используя оператор select.

	(Всего отправляется 5 значений, в один из свободных каналов)
*/
package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			select {
			case ch1 <- i:
				fmt.Println("Sent to ch1:", i)
			case ch2 <- i:
				fmt.Println("Sent to ch2:", i)
			}
		}
	}()

	for i := 0; i < 5; i++ {
		select {
		case val := <-ch1:
			fmt.Println("Received from ch1:", val)
		case val := <-ch2:
			fmt.Println("Received from ch2:", val)
		}
	}
}
