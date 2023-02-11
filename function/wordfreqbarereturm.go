/*
Упражнение 4.9. Напишите программу wordfreq для подсчета частоты каждого
слова во входном текстовом файле. Вызовите input.Split(bufio.ScanWords) до
первого вызова Scan для разбивки текста на слова, а не на строки.

Упражнение 5.5. Реализуйте функцию contWordsAndlmages(см. разделение
на слова в упр. 4.9).
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]
	if len(files) == 0 {
		println("need filepaht, exit")
		return
	}
	wrdcnt, err := countWords(files[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
	}
	if wrdcnt != nil {
		for w, c := range wrdcnt {
			fmt.Printf("%-25s\t%d\n", w, c)
		}
	}
}

func countWords(filepaht string) (wrdcnt map[string]int, err error) {
	wordsCount := make(map[string]int)
	file, err := os.OpenFile(filepaht, os.O_RDONLY, 0444)
	if err != nil {
		println(err)
		err = fmt.Errorf("wrong filepaht %s: %s", filepaht, err)
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
	wrdcnt = wordsCount
	return
}
