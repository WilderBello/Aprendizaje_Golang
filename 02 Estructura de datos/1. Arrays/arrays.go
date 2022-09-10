package main

import "fmt"

func main() {
	var numeros [5] int // Inicia en el indice cero, son inmutables, es estatico
	fmt.Println(numeros)

	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 30

	fmt.Println(numeros)
	fmt.Println(numeros[1])

	// Array con valores
	nombres := [3]string{"Wilder", "Ofrey", "Bello"}
	fmt.Println(nombres)

	// Array sin longitud definida
	colores := [...]string{"Rojo", "Verde", "Negro", "Azul"}
	fmt.Println(colores, len(colores)) // Array sin longitud definida y longitud

	// Indices definidos
	monedas := [...]string{0:"Dolar", 2:"Soles", 3:"Pesos", 5:"Euro"}
	fmt.Println(monedas, len(monedas))

	// Sub array
	subMoneda := monedas[:3]
	fmt.Println(subMoneda)
}