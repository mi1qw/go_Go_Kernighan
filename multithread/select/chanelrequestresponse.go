/*
использования select для записи в канал и ожидания
ответа из другого канала
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	request := make(chan string)
	response := make(chan string)

	go func() {
		for {
			req := <-request
			response <- "Received request: " + req
		}
	}()

	for i := 0; i < 6; i++ {
		select {
		case request <- "message-" + strconv.Itoa(i):
			fmt.Println("Sent message")
		case res := <-response:
			fmt.Println(res)
		}
	}
}
