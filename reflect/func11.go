package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type Mono[V comparable] struct {
	prev  *Mono[V]
	f     reflect.Value
	Value V
}

func Map1[T, V comparable](m *Mono[T], f func(T) V) *Mono[V] {
	// Создаем reflect.Value из функции
	fValue := reflect.ValueOf(f)
	newMono := &Mono[V]{
		prev: m,
		f:    fValue,
	}
	return newMono
}

func main() {
	// Пример использования Map1
	m := &Mono[int]{}
	f := func(i int) string {
		return strconv.Itoa(i*2) + " !"
	}
	result := Map1[int, string](m, f)

	// Вызов сохраненной функции
	args := []reflect.Value{reflect.ValueOf(5)} // Аргументы функции
	resultValue := result.f.Call(args)          // Вызов функции

	// Получение результата из resultValue
	if len(resultValue) > 0 {
		resultInt := resultValue[0].Interface().(string)
		fmt.Println(resultInt)
	}
}
