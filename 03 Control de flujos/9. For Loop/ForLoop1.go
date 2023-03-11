package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	nombre := "Eduardo"

	for i := 0; i < len(nombre); i++ {
		fmt.Println(string(nombre[i]))
	}

	numero := 0

	for numero != 1 {
		fmt.Println("Ejecutando For")
		numero = 1
	}
}
