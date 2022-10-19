package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func euler() {
	/*c := 3 + 4i
	fmt.Println(cmplx.Abs(c))*/
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)

	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
}

func triangle() {
	var a, b int = 3, 4
	c := calcTriangle(a, b)
	fmt.Println(c)
}

func calcTriangle(a int, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

func consts() {
	const filename string = "abc.txt"
	/*const a, b int = 3, 4*/
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enums() {
	/*const(
		cpp		= 0
		java	= 1
		python	= 2
		golang	= 3
	)*/
	const (
		cpp = iota
		_
		python
		golang
		javascript
	)

	//byte, kb, mb, gb, tb, pb
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp, javascript, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("Hello World")

	euler()
	triangle()
	consts()
	enums()
}
