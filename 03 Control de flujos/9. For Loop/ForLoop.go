package main

import "fmt"

func main() {

	// Bucle infinito
	/*
		for {
			fmt.Println("Bucle infinito")
		}
	*/
	// Bucle tipo while
	numeros := 12455
	c := 0
	for numeros > 0 {
		numeros /= 10
		c++
	}

	fmt.Println("Cantidad de digitos es:", c)
	// Bucle for
	num := 10
	for i := 0; i < num; i++ {
		fmt.Println("For normal")
	}
}
