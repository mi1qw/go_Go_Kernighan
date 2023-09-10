// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 142.

// The sum program demonstrates a variadic function.
package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

/*
type Number interface {
	int | []int
}

*/

// !+

func max[T constraints.Ordered](vals ...T) T {
	return find(func(x1, x2 T) bool { return x1 > x2 }, vals[:])
}

func min[T constraints.Ordered](vals ...T) T {
	return find(func(x1, x2 T) bool { return x1 < x2 }, vals[:])
}

func find[T constraints.Ordered](f func(x1, x2 T) bool, vals []T) T {
	if len(vals) == 1 {
		return vals[0]
	}
	res := vals[0]
	for _, val := range vals[1:] {
		if f(val, res) {
			res = val
		}
	}
	return res
}

// !-
func aaa[T constraints.Integer](a T) {
	fmt.Printf("%064[1]b\t%[1]d\n", ^T(0)>>1)
	fmt.Printf("%064[1]b\t%[1]d\n", ^T(0))

	fmt.Printf("%08[1]b\t%[1]d\n", ^T(1)<<1)
}
func main() {
	//aaa(0)
	//!+slice
	//values := []int{10, 20, 30, 40}
	//fmt.Println(max(values...)) // "10"

	println(max(1, 2, 3, 4))
	println(max(0))
	println(min(1, 2, 3, 4))
	println(min(1, 1, 1))
	//println(max([]int{}...))
	//!-slice
	println(" -------------- MyStruct --------------- ")

	ob := MyStruct[string]{Val: "Hello"}
	println(ob.GetVal())
	fmt.Printf("GetMystruct = %T \n\n", ob.GetMystruct())

	data := Data{
		name: "Data String",
	}
	obData := MyStruct[Data]{
		Val: data,
	}
	//println(obData.GetVal())
	fmt.Printf("%s \n", obData.GetVal().name)
	fmt.Printf("GetMystruct = %T \n\n", obData.GetMystruct())

	myStructData := NewMyStruct[Data]()
	myStructData.Val = data
	fmt.Printf("%s \n", obData.GetVal().name)

}

type MyStruct[T any] struct {
	Val T
}

func (m *MyStruct[T]) GetVal() T {
	return m.Val
}
func (m *MyStruct[T]) GetMystruct() *MyStruct[T] {
	return m
}

func NewMyStruct[T any]() *MyStruct[T] {
	return &MyStruct[T]{}
}

type Data struct {
	name string
}
