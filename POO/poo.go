package main

import "fmt"

type persona struct {
	nombre   string
	apellido string
	edad     int
}

func main() {
	persona1 := persona{"Wilder", "Bello", 26}
	fmt.Println(persona1)

	persona2 := persona{"Paula", "Bello", 29}
	fmt.Println(persona2)

	fmt.Println(persona1.nombre, persona2.nombre)

	persona1.edad = 27
	fmt.Println(persona1)
}
