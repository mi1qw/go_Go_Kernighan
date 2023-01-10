/*
Упражнение 4.6. Напишите функцию, которая без выделения дополнительной па
мяти преобразует последовательности смежных пробельных символов Unicode (см.
Unicode.IsSpace) в срезе []byte в кодировке UTF-8 в один пробел ASCII.
*/
package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	s := "Hello,    Г 1  ФЯя  d"
	fmt.Printf("%s\nДлина строки(байт/ascii) %d\n"+
		"Колличество рун UTF-8    %d\n\n",
		s, len(s), utf8.RuneCountInString(s))
	//printRunes(s)
	s, newSize := nonspace(s)
	fmt.Printf("%s\tновый размер-%d рун", s, newSize)
}

func nonspace(strings string) (string, int) {
	bytes := []byte(strings)
	var isprevSpace, isdouble bool
	var i, b, bDelta int
	for n := utf8.RuneCount(bytes); n > 0; n-- {
		r, size := utf8.DecodeRune(bytes[b:])
		//fmt.Printf("%c %v \n", r, unicode.IsSpace(r))
		if !unicode.IsSpace(r) {
			isprevSpace = false
			copy(bytes[i:], bytes[b:])
			if isdouble {
				b -= bDelta
				bDelta = 0
				isdouble = false
			}
			i += size
		} else {
			if isprevSpace {
				bDelta++
				isdouble = true
			} else {
				i += size
			}
			isprevSpace = true
		}
		b += size
		//	fmt.Println(string(bytes[:i]))
		//	fmt.Println(bytes[:])
	}
	return string(bytes[:i]), utf8.RuneCount(bytes[:i])
}

func printRunes(s string) {
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\t%d\n", i, r, size)
		i += size
	}
}
