package main

import "fmt"

type persona struct {
	nombre   string
	apellido string
	edad     int
}

func (p persona) saludar(saludo string) {
	fmt.Println(saludo + ", " + p.nombre)
}

func (pers persona) cumple() int {
	return pers.edad + 1
}

func main() {
	persona1 := persona{"Wilder", "Bello", 26}

	persona1.saludar("Hola")

	edad_persona1 := persona1.cumple()
	fmt.Println(edad_persona1)
}
