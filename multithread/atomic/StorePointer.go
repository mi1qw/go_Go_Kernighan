package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

func main() {
	var activeCh unsafe.Pointer
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)
	atomic.StorePointer(&activeCh, unsafe.Pointer(&ch1)) // Начинаем с ch1 как активного канала

	var wg sync.WaitGroup
	wg.Add(1) // Ожидаем две горутины

	// Горутина, которая периодически меняет активный канал
	go func() {
		//defer wg.Done()
		var prev unsafe.Pointer
		var next unsafe.Pointer = unsafe.Pointer(&ch2)
		for {
			// Добавьте небольшую задержку, чтобы не нагружать процессор
			//time.Sleep(time.Second)
			prev = atomic.SwapPointer(&activeCh, next)
			next = atomic.SwapPointer(&activeCh, prev)
		}
	}()

	var sumData int
	// Горутина, которая отправляет данные в активный канал
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			ch := *(*chan int)(atomic.LoadPointer(&activeCh)) // Получаем текущий активный канал
			ch <- i
			sumData += i
		}
		//close(*(*chan int)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&ch1)))))
	}()

	var sumGetData int
	// Горутина, которая читает данные из активного канала
	go func() {
		for {

			ch := *(*chan int)(atomic.LoadPointer(&activeCh)) // Получаем текущий активный канал
			if len(ch) == 0 {
				continue
			}
			select {
			case val := <-ch:
				fmt.Printf("Received: %d  len: %d\n", val, len(ch))
				sumGetData += val
			}
		}
	}()

	// Ожидаем завершения обеих горутин
	wg.Wait()

	// Дайте программе время на выполнение
	time.Sleep(1 * time.Second)
	fmt.Printf("sumGetData = %d\n", sumGetData)
	fmt.Printf("sumData = %d\n", sumData)
}
