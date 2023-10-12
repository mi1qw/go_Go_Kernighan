package main

import (
	"fmt"
	"reflect"
)

type MonoF[V comparable] struct {
	prev *MonoF[V]
	f    reflect.Value
}

func MapF[V comparable](m *MonoF[V], f interface{}) *MonoF[V] {
	// Создаем reflect.Value из функции
	fValue := reflect.ValueOf(f)
	// Проверяем, что fValue - это функция
	if fValue.Kind() != reflect.Func {
		panic("f is not a function")
	}

	return &MonoF[V]{
		prev: m,
		f:    fValue,
	}
}

func main() {
	m := &MonoF[int]{}
	f := func(i int) int {
		return i * 2
	}
	result := MapF(m, f)

	// Вызов сохраненной функции
	args := []reflect.Value{reflect.ValueOf(5)} // Аргументы функции
	resultValue := result.f.Call(args)          // Вызов функции

	// Получение результата из resultValue
	if len(resultValue) > 0 {
		resultInt := resultValue[0].Interface().(int)
		fmt.Println(resultInt)
	}
}
