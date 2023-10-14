package main

import (
	"fmt"
)

type Printer[T any] interface {
	//PrintValue(value T)
	Print(int)
}

type MyPrinter[T any] struct{}

func (p MyPrinter[T]) PrintValue(value T) {
	fmt.Println(value)
}
func (p MyPrinter[T]) Print(value int) {
	fmt.Println(value)
}

type AnyPrinter[T any] struct{}

func (p AnyPrinter[T]) PrintValue(value T) {
	fmt.Println(value)
}
func (p AnyPrinter[T]) Print(value int) {
	fmt.Println(value)
}
func main() {
	sl := []Printer[int]{MyPrinter[int]{}, AnyPrinter[int]{}}
	for i, p := range sl {
		//p.PrintValue(i)
		p.Print(i)
	}
}
