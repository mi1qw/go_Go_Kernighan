// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 142.

// The sum program demonstrates a variadic function.
package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"reflect"
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

type AnyStruct[T any] struct {
	val  any
	Tval T
}

func (s *AnyStruct[T]) PutT(str any) {
	t := str.(T)
	s.val = t
	s.Tval = t
}
func init() {
	println("start !")
}
func main() {
	println(" -------------- Holder any --------------- ")
	anystruct := AnyStruct[Data]{}
	anystruct.PutT(Data{
		name: "Data String",
	})
	fmt.Printf("GetMystruct = %T \n\n", anystruct)
	v1 := anystruct.Tval.GetName()
	v2 := anystruct.val.(Data).GetName()
	println(v1, v2)

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
	fmt.Printf("%s \n\n", obData.GetVal().name)

	println(" -------------- MyStruct any --------------- ")
	holder := Holder{
		builder: myStructData,
	}
	fmt.Printf("%t \n", holder)
	fmt.Printf("%T \n\n", holder.builder)
	// приведение к типу
	myStruct := holder.builder.(*MyStruct[Data])
	fmt.Printf("%t \n", myStruct)
	fmt.Printf("%T \n\n", myStruct)
	println(myStruct.Val.name)

	println(" -------------- Holder any --------------- ")
	h := Holder{
		builder:    data,
		hInterface: data,
	}
	//через интерфейс
	println(h.hInterface.GetName())

	get := h.get()
	println(get.name)

	targetType := reflect.TypeOf(Data{})
	value, err := h.getRef(targetType)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Теперь мы получили значение с нужным типом
	if d, ok := value.(Data); ok {
		fmt.Println(d.name) // Выведет "Hello, World!"
	} else {
		fmt.Println("Не удалось выполнить приведение типа.")
	}
}

type DataInterface interface {
	GetName() string
}

type Data struct {
	name string
}

func (d Data) GetName() string {
	return d.name
}

type Holder struct {
	builder    any
	typeRef    reflect.Type
	hInterface DataInterface
}

func (h Holder) get() Data {
	data := h.builder.(Data)
	return data
}

func (h Holder) getRef(targetType reflect.Type) (interface{}, error) {
	if reflect.TypeOf(h.builder) != targetType {
		return nil, fmt.Errorf("Type assertion failed. Expected type: %s", targetType)
	}
	return h.builder, nil
}
