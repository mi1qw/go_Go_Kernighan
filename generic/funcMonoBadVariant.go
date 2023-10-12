package main

import (
	"fmt"
	"strconv"
)

type Inter interface {
	Screen111() any
	getch() chan any
	getF() func(any) any
	setF(f func(any) any)
}
type MonoB[V comparable] struct {
	prev  Inter
	f     func(any) any
	Value V
	ch    chan any
}

func (m MonoB[V]) Screen111() any {
	// todo добавить wrong && m.Value == nil
	if m.f != nil {
		newVal := m.f(m.Value)
		m.ch <- newVal
	}

	prev := m.prev
	if prev != nil {
		select {
		case val := <-prev.getch():
			if v, ok := val.(V); ok {
				fmt.Printf("%v  value from prev chan \n", v)
				m.Value = v
			} else {
				println("error conversion:")

			}
		default:
		}
	}
	println(m.Value, " Screen")
	return m.Value
}
func (m MonoB[V]) getch() chan any {
	return m.ch
}
func (m MonoB[V]) getF() func(any) any {
	return m.f
}
func (m MonoB[V]) setF(f func(any) any) {
	m.f = f
}
func Map1[T, V comparable](mono *MonoB[T], f func(T) V) *MonoB[V] {
	mono.f = func(a any) any {
		t := a.(T)
		return f(t)
	}
	newMono := &MonoB[V]{
		prev: mono,
		ch:   make(chan any, 1),
	}
	return newMono
}

func main() {

	m := MonoB[int]{
		Value: 1,
		ch:    make(chan any, 1),
	}

	result := Map1[int, string](&m, func(i int) string {
		return strconv.Itoa(i*2) + " !"
	})

	result.prev.Screen111()
	result.Screen111()
}
