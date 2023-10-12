package main

import (
	"log"
	"strconv"
)

type Connector interface {
	ComputeVal(func(any) any) any
	getch() chan any
}

type Mono[V comparable] struct {
	prev  *Connector
	f     func(any) any
	Value V
	ch    chan any
}

func (m *Mono[V]) ComputeVal(f func(any) any) any {
	newVal := f(m.Value)
	//m.ch <- newVal
	return newVal
}

func (m *Mono[V]) getch() chan any {
	return m.ch
}

func (m *Mono[V]) Calculate() {
	newVal := (*m.prev).ComputeVal(m.f)
	if v, ok := newVal.(V); ok {
		m.Value = v
	} else {
		log.Panic("error conversion")
	}
}

func Map2[T, V comparable](mono *Mono[T], f func(T) V) *Mono[V] {
	fu := func(a any) any {
		t := a.(T)
		return f(t)
	}
	connector := Connector(mono)
	return &Mono[V]{
		prev: &connector,
		ch:   make(chan any, 1),
		f:    fu,
	}
}

// connector, ok := mono.(*Connector)
func main() {
	m := &Mono[int]{
		Value: 1,
		ch:    make(chan any, 1),
	}
	result := Map2[int, string](m, func(i int) string {
		return strconv.Itoa(i*2) + " !"
	})
	result.Calculate()
	println("Calculated val  ", result.Value)
}
