package main

type item[T any] struct {
	value T
	next  *item[T]
	prev  *item[T]
}

type LinkedList[T any] struct {
	//head *item[T]
	tail *item[T]
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{
		//head: nil,
		tail: nil,
	}
}

func (i *LinkedList[T]) add(t T) {
	node := &item[T]{
		value: t,
		prev:  i.tail,
	}
	if i.tail == nil {
		i.tail = node
	} else {
		i.tail.next = node
		i.tail = node
	}
}

func (i *LinkedList[T]) removeLast() T {
	r := i.tail.value
	i.tail = i.tail.prev
	i.tail.next = nil
	return r
}

func (i *LinkedList[T]) isEmpty() bool {
	return i.tail == nil
}

func (i *LinkedList[T]) forEach(f func(t T)) {
	if i.isEmpty() {
		return
	}
	var head = i.tail
	for head.prev != nil {
		head = head.prev
	}
	for {
		f(head.value)
		if head.next == nil {
			break
		}
		head = head.next
	}
}
