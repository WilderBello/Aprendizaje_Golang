package main

import "fmt"

func main() {
	a := 4
	b := 5

	var r bool

	// igualdad
	r = a == b
	fmt.Printf("%d es igual a %d? %t \n", a, b, r)

	// distinto
	r = a != b
	fmt.Printf("%d es distinto a %d? %t \n", a, b, r)

	// Mayor que
	r = a > b
	fmt.Printf("%d es mayor a %d? %t \n", a, b, r)

	// Menor que
	r = a < b
	fmt.Printf("%d es menor a %d? %t \n", a, b, r)

	// mayor o igual
	r = a >= b
	fmt.Printf("%d es mayor o igual a %d? %t \n", a, b, r)

	// menor o igual
	r = a <= b
	fmt.Printf("%d es menor o igual a %d? %t \n", a, b, r)

}
