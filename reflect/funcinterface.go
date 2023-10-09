package main

import (
	"fmt"
	"reflect"
)

type Mono[V comparable] struct {
	prev *Mono[V]
	f    reflect.Value
}

func Map1[V comparable](m *Mono[V], f interface{}) *Mono[V] {
	// Создаем reflect.Value из функции
	fValue := reflect.ValueOf(f)
	reflect.TypeOf(V)
	// Проверяем, что fValue - это функция
	if fValue.Kind() != reflect.Func {
		panic("f is not a function")
	}

	newMono := &Mono[V]{
		prev: m,
		f:    fValue,
	}
	return newMono
}

func main() {
	// Пример использования Map1
	m := &Mono[int]{}
	f := func(i int) int {
		return i * 2
	}
	result := Map1(m, f)

	// Вызов сохраненной функции
	args := []reflect.Value{reflect.ValueOf(5)} // Аргументы функции
	resultValue := result.f.Call(args)          // Вызов функции

	// Получение результата из resultValue
	if len(resultValue) > 0 {
		resultInt := resultValue[0].Interface().(int)
		fmt.Println(resultInt)
	}
}
