package main

import "fmt"

func main() {
	hola := "Hola"
	mundo := "Mundo"

	fmt.Println(hola)
	fmt.Println(mundo)

	nombre := "Wilder"
	edad := 26

	fmt.Printf("Hola, %s y su edad es %d\n", nombre, edad)  // %s -> string,  %d -> numero entero, %v -> no sabemos el tipo de dato.

	fmt.Printf("Hola, %v y su edad es %v \n", nombre, edad)

	// Guardar en variable
	mensaje := fmt.Sprintf("Hola, %v y su edad es %v", nombre, edad)

	fmt.Println(mensaje)

	// Tipo de dato de una variable: %T
	fmt.Printf("nombre: %T \n", nombre)

	// Recibir datos por consola
	var apellido string
	fmt.Print("Ingrese otro apellido: ")
	fmt.Scanln(&apellido)

	fmt.Print("Otro apellido: ", apellido)
}