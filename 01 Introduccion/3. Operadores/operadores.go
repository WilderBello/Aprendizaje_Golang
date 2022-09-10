package main

import "fmt"

/* 
Comentario de bloque
*/
func main() {
	a := 20
	b := 10

	//Suma
	result := a + b
	fmt.Println("Suma: ", result)
	
	result = a - b
	fmt.Println("Resta: ", result)

	result = a * b
	fmt.Println("Multi: ", result)

	result = a / b
	fmt.Println("Div: ", result)

	var div float64 = 3.0 / 2.0 //Devuelve el entero
	fmt.Println("Div decimal: ", div)

	result = 3 % 2
	fmt.Println("Modulo: ", result)
}