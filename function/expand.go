package main

import (
	"fmt"
	"strings"
)

// У пражнение 5.9. Напишите функцию expand(s string, f func(string) string) string
// которая заменяет каждую подстроку "$foo” в s текстом, который
// возвращается вызовом f ("foo").

func main() {
	s := expand("Hello $foo World", funcString)
	fmt.Printf("%s", s)
}

func expand(s string, f func(str string) string) string {
	return strings.Replace(s, "$foo", f("foo"), -1)
}

func funcString(str string) string {
	return str
}
