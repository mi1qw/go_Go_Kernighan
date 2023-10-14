package main

import "fmt"

type MyT int

/*
храним в битах
*/
const (
	VAR1 int = 1 << iota
	VAR2
	VAR3
	VAR4
	VAR5
	VAR6
	VAR7
	VAR8
	VAR9
	VAR10
)

func main() {
	println(VAR1, VAR2)
	fmt.Printf("%b\n", VAR1)
	fmt.Printf("%b\n", VAR2)
	fmt.Printf("%b\n", VAR3)
	fmt.Printf("%b\n", VAR4)
	fmt.Printf("%b\n", VAR5)
	fmt.Printf("%b\n", VAR6)
	fmt.Printf("%b\n", VAR7)
	fmt.Printf("%b\n", VAR8)
	fmt.Printf("%b\n", VAR9)
	fmt.Printf("%b\n", VAR10)

	v := VAR10 | VAR1
	fmt.Printf("%b\n", v)
	fmt.Printf("%b\n", v&VAR10)

}
