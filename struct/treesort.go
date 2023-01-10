package main

import (
	"fmt"
	"strings"
)

type tree struct {
	User
	left, right *tree
}
type User struct {
	name  string
	value int
}

func (u User) less(u2 User) bool {
	if strings.Compare(u.name, u2.name) == -1 {
		return true
	}
	return false
}

func main() {
	users := []User{
		{"b", 20},
		{"e", 25},
		{"d", 50},
		{"a", 10},
		{"c", 30},
	}
	fmt.Printf("%v\n", users)
	Sort(users)
	fmt.Printf("%v\n", users)
}

// Sort сортирует значения "на месте",
func Sort(values []User) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

func add(t *tree, value User) *tree {
	if t == nil {
		t = new(tree)
		t.User = value
		return t
	}
	if value.less(t.User) {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

// appendValues добавляет элементы t к values в требуемом
// порядке и возвращает результирующий срез,
func appendValues(values []User, t *tree) []User {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.User)
		values = appendValues(values, t.right)
	}
	return values
}
