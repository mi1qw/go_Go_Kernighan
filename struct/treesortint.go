package main

import "fmt"

type treeint struct {
	value       int
	left, right *treeint
}

func main() {
	v := []int{5, 3, 4, 7}
	fmt.Printf("%v\n", v)
	SortInt(v)
	fmt.Printf("%v\n", v)
}

// SortInt сортирует значения "на месте",
func SortInt(values []int) {
	var root *treeint
	for _, v := range values {
		root = addInt(root, v)
	}
	appendValuesInt(values[:0], root)
}

// appendValuesInt добавляет элементы t к values в требуемом
// порядке и возвращает результирующий срез,
func appendValuesInt(values []int, t *treeint) []int {
	if t != nil {
		values = appendValuesInt(values, t.left)
		values = append(values, t.value)
		values = appendValuesInt(values, t.right)
	}
	return values
}
func addInt(t *treeint, value int) *treeint {
	if t == nil {
		// Эквивалентно возврату &treeint{value: value},
		t = new(treeint)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = addInt(t.left, value)
	} else {
		t.right = addInt(t.right, value)
	}
	return t
}
