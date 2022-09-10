package main

import "fmt"

func main() {
	var nombre1 string
	nombre1 = "Wilder"

	var nombre2 string = "Ofrey"

	edad := 26 // Estamos definiendo la variable directamente y por lo tanto se establece el tipo de dato.

	fmt.Println(nombre1, nombre2, edad)

	// Las variables traen un valor por defecto
	var a int
	var b float64
	var c string
	var d bool
	fmt.Println(a, b, c,d)

	const pi = 3.141592653589
	fmt.Println(pi)
}