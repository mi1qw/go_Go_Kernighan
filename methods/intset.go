// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 165.

/*
Упражнение 6.1. Реализуйте следующие дополнительные методы:
func (*IntSet) Len() int 		// Возвращает количество элементов
func (*IntSet) Remove(x int) 	// Удаляет x из множества
func (*IntSet) Clear() 			// Удаляет все элементы множества
func (*IntSet) Copy() *IntSet 	// Возвращает копию множества

Упражнение 6.2. Определите вариативный метод (*IntSet).AddAll(...int),
который позволяет добавлять список значений, например s.AddAll(1,2,3).

Упражнение 6.3. (*IntSet).UnionWith вычисляет объединение двух мно
жеств с помощью оператора |, побитового оператора ИЛИ. Реализуйте методы
IntersectWith, DifferenceWith и SymmetricDifference для соответствующих
операций над множествами. (Симметричная разность двух множеств содержит эле
менты, имеющиеся в одном из множеств, но не в обоих одновременно.)

*/

// Package intset provides a set of integers based on a bit vector.
package main

import (
	"bytes"
	"fmt"
)

//!+intset

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

func main() {
	var x, y IntSet

	x.Add(1)
	x.Add(8)
	x.Add(40)
	println("x", x.String(), x.Len())

	a := x.Copy()

	x.Remove(8)
	println("x remove", x.String(), x.Len())

	x.Clear()
	println("x", x.String(), x.Len())
	println("Copy", a.String(), a.Len())

	y.AddAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 255)
	println("y", y.String(), y.Len())

	// IntersectWith
	for i := 0; i < 8; i++ {
		x.Add(i)
		// fmt.Printf("%08b\t%d\n", x.words, i)
	}
	x.Remove(5)
	println("x", x.String(), x.Len())
	// x.AddAll(65)
	fmt.Printf("%08b x\n", x.words)

	y.Clear()
	y.AddAll(0, 2, 3, 4, 5, 64)

	fmt.Printf("%08b y\n", y.words)
	fmt.Printf("%08b Intersect\n", x.IntersectWith(&y).words)
	fmt.Printf("%08b Intersect\n", y.IntersectWith(&x).words)
	fmt.Printf("%08b Difference\n", x.DifferenceWith(&y).words)
	fmt.Printf("%08b Difference\n", y.DifferenceWith(&x).words)
	fmt.Printf("%08b x SymmetricDifference y \n", x.SymmetricDifference(&y).words)
	fmt.Printf("%08b y SymmetricDifference x \n", y.SymmetricDifference(&x).words)
}

func (s *IntSet) AddAll(n ...int) {
	for _, x := range n {
		s.Add(x)
	}
}

func (s *IntSet) Copy() *IntSet {
	var x IntSet
	x.words = make([]uint64, len(s.words))
	copy(x.words, s.words)
	return &x
}

func (s *IntSet) Clear() {
	s.words = make([]uint64, 0)
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if word < len(s.words) {
		s.words[word] &^= 1 << bit
	}
}

func (s *IntSet) Len() int {
	count := 0
	for _, word := range s.words {
		for i := 0; i < 64; i++ {
			if word&(1<<uint(i)) != 0 {
				count++
			}
		}
	}
	return count
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func copyApply(a, b *IntSet, f func(uint64, uint64) uint64) *IntSet {
	var res IntSet
	for i, word := range a.words {
		res.words = append(res.words, f(word, b.words[i]))
	}
	return &res
}

func (s *IntSet) IntersectWith(t *IntSet) *IntSet {
	lens, lent := len(s.words), len(t.words)
	if lent >= lens {
		return copyApply(s, t, func(a, b uint64) uint64 { return a & b })
	}
	return copyApply(t, s, func(a, b uint64) uint64 { return a & b })
}

func (s *IntSet) DifferenceWith(t *IntSet) *IntSet {
	lens, lent := len(s.words), len(t.words)
	if lent >= lens {
		diff := copyApply(s, t, func(a, b uint64) uint64 { return a ^ b })
		diff.words = append(diff.words, t.words[lens:lent]...)
		return diff
	}
	diff := copyApply(t, s, func(a, b uint64) uint64 { return a ^ b })
	diff.words = append(diff.words, s.words[lent:lens]...)
	return diff
}

func (s *IntSet) SymmetricDifference(t *IntSet) *IntSet {
	lens, lent := len(s.words), len(t.words)
	if lens > lent {
		diff := copyApply(t, s, func(b, a uint64) uint64 { return a &^ b })
		diff.words = append(diff.words, s.words[lent:lens]...)
		return diff
	}
	return copyApply(s, t, func(a, b uint64) uint64 { return a &^ b })
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
