/*
использования select для таймаута операции на канале
Канал ch ждёт ответа в течении  5 секунд, а иначе "Timeout"
Для примера посылаем int в канал через 2 секкунды,
а значит Timeout непроизойдёт

В Golang <-chan T и chan T - это два разных типа канала,
где <-chan T является типом только для чтения (т.е. можно только читать значения из канала),
а chan T - только для записи (т.е. можно только записывать значения в канал).
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	var timeChan <-chan time.Time                 // переменная типа <-chan time.Time
	timeChan = time.After(100 * time.Millisecond) // возвращает канал для чтения, не пауза!
	fmt.Printf("%T - timeout\n ", timeChan)

	<-time.After(1 * time.Second) // прочитать из канала, поэтому работает как пауза
	fmt.Println("Прошла 1 секунда")

	fmt.Println("Старт")

	// отправляем значение в ожидающий канал
	go func() {
		time.Sleep(2 * time.Second)
		ch <- 123 // отправляем значение в канал, закомментировать
	}()

	select {
	case val := <-ch:
		fmt.Println("Received:", val)
	case <-time.After(5 * time.Second): // пытаемя прочитать из канала после 5 сек.
		fmt.Println("Timeout")
	}
}
