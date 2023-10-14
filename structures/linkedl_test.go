package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewLinkedList(t *testing.T) {
	list := NewLinkedList[int]()
	list.add(1)
	list.add(2)
	list.add(3)
	list.forEach(func(n int) {
		t.Log(n)
	})
}

func TestNewLinkedListRemove(t *testing.T) {
	list := NewLinkedList[int]()
	list.add(1)
	list.add(2)
	list.add(3)
	removed := list.removeLast()
	t.Logf("removed %d\n", removed)
	println()
	list.forEach(func(n int) {
		t.Log(n)
	})
	assert.Equal(t, 3, removed)
}
