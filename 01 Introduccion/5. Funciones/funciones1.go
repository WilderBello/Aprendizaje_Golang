package main

import "fmt"

func mi_func() {
	fmt.Println("Mi Funcion")
}

func saludar(nombre string) {
	fmt.Println("Hola " + nombre)
}

func sumar(num1 int, num2 int) int {
	suma := num1 + num2
	return suma
}

func main() {
	mi_func()
	saludar("Wilder")
	fmt.Println(sumar(1, 2))
}
