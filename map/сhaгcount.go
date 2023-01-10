/*
Упражнение 4.8. Измените сhaгcount так, чтобы программа подсчитывала коли
чество букв, цифр и прочих категорий Unicode с использованием функций наподобие
Unicode.IsLetter.
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	/*
		seen := make(map[string]bool) // Множество строк
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			line := input.Text()
			if !seen[line] {
				seen[line] = true
				fmt.Println(line)
			}
		}
		if err := input.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
			os.Exit(1)
		}
	*/

	counts := make(map[rune]int)    // Количество символов Unicode
	var utflen [utf8.UTFMax + 1]int // Количество длин кодировок UTF-8
	invalid := 0                    // Количество некорректных символов UTF-8
	typeRune := map[string]int{
		"letter":      0,
		"mark":        0,
		"number":      0,
		"punctuation": 0,
		"space":       0,
		"symbolic":    0,
	}
	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // Возврат руны, байтов, ошибки
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
		//fmt.Printf("%d -\n", r)
		if r == 48 {
			break
		}
		if unicode.IsSpace(r) {
			typeRune["space"]++
		} else if unicode.IsLetter(r) {
			typeRune["letter"]++
		} else if unicode.IsMark(r) {
			typeRune["mark"]++
		} else if unicode.IsNumber(r) {
			typeRune["number"]++
		} else if unicode.IsPunct(r) {
			typeRune["punctuation"]++
		} else if unicode.IsSymbol(r) {
			typeRune["symbolic"]++
		}
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d неверных символов UTF-8\n", invalid)
	}
	fmt.Print("\ntypeRune\tcount\n")
	for t, n := range typeRune {
		fmt.Printf("%-12s\t%d\n", t, n)
	}
}
