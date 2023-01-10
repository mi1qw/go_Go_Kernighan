/*
Упражнение 4.9. Напишите программу wordfreq для подсчета частоты каждого
слова во входном текстовом файле. Вызовите input.Split(bufio.ScanWords) до
первого вызова Scan для разбивки текста на слова, а не на строки.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	wordsCount := make(map[string]int)
	files := os.Args[1:]
	fmt.Printf("%s", files)
	if len(files) == 0 {
		println("need filepaht, exit")
		return
	}
	file, err := os.OpenFile(files[0], os.O_RDONLY, 0444)
	if err != nil {
		println(err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			println(err)
		}
	}(file)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		text := scanner.Text()
		wordsCount[text]++
	}
	if err = scanner.Err(); err != nil {
		println(err)
	}
	for w, c := range wordsCount {
		fmt.Printf("%-19s\t%d\n", w, c)
	}
}
