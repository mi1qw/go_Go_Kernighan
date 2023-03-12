package main

/*
У пражнение 7.2. Напишите функцию CountingWriter с приведенной ниже сиг
натурой, которая для данного io.Writer возвращает новый Writer, являющийся
оболочкой исходного, и указатель на переменную int64, которая в любой момент
содержит количество байтов, записанных в новый Writer.
func CountingWriter(w io.Writer) (io.Writer, *int64)
*/
import (
	"fmt"
	"io"
	"os"
)

type Count struct {
	w io.Writer
	c int64
}

func (c *Count) Write(p []byte) (int, error) {
	n, err := c.w.Write(p)
	if err != nil {
		return 0, err
	}
	c.c += int64(n)
	return n, nil
}
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	counter := Count{w, 0}
	return &counter, &counter.c
}

func main() {
	writer, c := CountingWriter(os.Stdout)
	println(*c, "bytes")
	fmt.Fprintf(writer, "Hello, %s!!!\n", "world")
	println(*c, "bytes")
	fmt.Fprintf(writer, "one more time\n")
	println(*c, "bytes")
}
