/*
Упражнение 4.7. Перепишите функцию reverse так, чтобы она без выделения
дополнительной памяти обращала последовательность символов среза []byte , кото
рый представляет строку в кодировке UTF-8. Сможете ли вы обойтись без выделения
новой памяти?
*/
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//str := "abcабс"
	str := "гпcYнqы"
	a := []byte(str)
	fmt.Printf("%s\n", a)
	reverseStringUTF(&a)
	fmt.Printf("%s\n", a)
}

func reverseStringUTF(bytes1 *[]byte) {
	var bytes = *bytes1
	var sizeF, sizeL, sizeM int
	var bF = make([]byte, len(bytes))
	var bFi, bLi int
	bLi = len(bytes)
	runeCount := utf8.RuneCount(bytes)
	f, l := 0, len(bytes)
	for loop := runeCount / 2; loop > 0; loop-- {
		_, sizeF = utf8.DecodeRune(bytes[f:])
		_, sizeL = utf8.DecodeLastRune(bytes[:l])
		copy(bF[bFi:], bytes[l-sizeL:l])
		bFi += sizeL
		copy(bF[bLi-sizeF:], bytes[f:f+sizeF])
		bLi -= sizeF
		f, l = f+sizeF, l-sizeL
	}
	if runeCount%2 != 0 {
		_, sizeM = utf8.DecodeRune(bytes[f:])
		copy(bF[bFi:], bytes[f:f+sizeM])
	}
	copy(*bytes1, bF)
}
