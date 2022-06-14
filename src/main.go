package main

import "fmt"

func main() {
	const pi float64 = 3.14
	const pi2 = 3.1416

	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	// Variables
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	// Zero Values
	var a int
	var b string
	var c float64
	var d bool

	fmt.Println(a, b, c, d)

	// Calc square area

	const squareBase = 10
	squareArea := squareBase * squareBase

	fmt.Println("Square Area:", squareArea)
}
