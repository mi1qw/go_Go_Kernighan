/*
Упражнение 4.5. Напишите функцию, которая без выделения дополнительной па
мяти удаляет все смежные дубликаты в срезе []string.
*/
package main

import "fmt"

func main() {
	data := []string{"one", "one", "two", "three"}
	fmt.Println(nondouble(data))
}
func nondouble(strings []string) []string {
	i := 0
	var prev string
	for _, s := range strings {
		if s != prev {
			strings[i] = s
			i++
		}
		prev = s
	}
	return strings[:i]
}
