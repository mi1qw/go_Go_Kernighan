package main

import (
	"strconv"
)

type Mono[T, V comparable] struct {
	prev  *Mono[V, T]
	f     func(T) V
	Value V
}

func (m *Mono[T, V]) map2(f func(V) T) *Mono[V, T] {
	newMono := &Mono[V, T]{
		prev: m,
		f:    f,
		//Value: zeroValue(V), // Здесь можете установить начальное значение Value, если нужно
	}
	return newMono
}

//func Map1[T, V comparable](m *Mono[T, V], f func(T) V) *Mono[T, V] {
//	newMono := &Mono[T, V]{
//		prev: m,
//		f:    f,
//		//Value: zeroValue(V), // Здесь можете установить начальное значение Value, если нужно
//	}
//	return newMono
//}

func main() {
	// Пример использования Map1
	m := &Mono[int, int]{Value: 20}
	f := func(i int) string {
		return strconv.Itoa(i*2) + " !"
	}
	println(m, f)

	map21 := m.map2(f)
	println(map21.Value)
	//result := Map1(m, f)
	//println(result)
	//
	//// Вызов сохраненной функции
	//result.Value = result.f(5) // Вызов функции и установка результата в Value
	//
	//// Получение результата
	//fmt.Println(result.Value)
}
