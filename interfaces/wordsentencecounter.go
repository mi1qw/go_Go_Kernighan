package main

/*
У пражнение 7.1. Используя идеи из ByteCounter, реализуйте счетчики для слов
и строк. Вам пригодится функция bufio.ScanWords.

*/
import (
	"bufio"
	"fmt"
	"os"
)

type WordSentCounter struct {
	words    int
	sentence int
}

func (w *WordSentCounter) Write(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}
	w.sentence++
	w.words++
	for _, b := range p {
		if b == ' ' {
			w.words++
		}
	}
	return len(p), nil
}

func (w *WordSentCounter) String() string {
	return fmt.Sprintf("\nwords %d\tsentence %d\n", w.words, w.sentence)
}

func main() {
	files := os.Args[1:]
	if len(files) == 0 {
		println("need filepaht, exit")
		return
	}
	fmt.Printf("%s", files[0])
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
	counter := WordSentCounter{0, 0}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Fprintf(&counter, "%s", text)
	}
	if err = scanner.Err(); err != nil {
		println(err)
	}
	println(counter.String())
}
