package main

import "fmt"

func main() {
	// Mapas es parecido a los diccionarios de python.
	dias := make(map[int]string)

	fmt.Println(dias)

	//Agregar datos clave-valor
	dias[1] = "Domingo"
	dias[2] = "Lunes"
	dias[3] = "Martes"
	dias[4] = "Miercoles"
	dias[5] = "Jueves"
	dias[6] = "Viernes"
	dias[7] = "Sabado"

	fmt.Println(dias)

	// Modificar datos clave-valor
	dias[7] = "SABADO"
	fmt.Println(dias)

	//Eliminar datos por clave
	delete(dias, 1)
	fmt.Println(dias)

	// Nuevo Mapa
	estudiantes := make(map[string][]int)

	estudiantes["Wilder"] = []int{13, 15, 16}
	estudiantes["Bello"] = []int{14, 13, 17}
	
	fmt.Println(estudiantes)

	// Acceder a elementos
	fmt.Println(estudiantes["Wilder"][1])
}