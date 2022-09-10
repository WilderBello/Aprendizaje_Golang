package main

import "fmt"

func saludar(nombre string) {
	//fmt.Println("Hola mundo")
	fmt.Println("Hola,", nombre)
}

func sumar(a, b int) int { 
	// Debe especificar el tipo de dato a retornar dentro de la definicion.
	return a + b
}

func main() {
	saludar("Wilder")
	r := sumar(1,3)

	fmt.Println("La suma es:", r)
}