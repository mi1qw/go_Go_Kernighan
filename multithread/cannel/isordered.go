package main

func main() {
	ints := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ints <- i
	}
	for i := 0; i < 10; i++ {
		n := <-ints
		println(n)
	}
}
