package main

import "fmt"

func main() {
	if nombre, edad := "Alex", 26; nombre == "Alex" {
		fmt.Println("Hola, ", nombre)
	} else {
		fmt.Printf("Nombre: %s - Edad: %d \n", nombre, edad)
	}

	// Obtener valor de mapa
	users := make(map[string]string)

	users["Alex"] = "alex@gmail.com"
	users["Roel"] = "roel@gmail.com"

	//correo, ok := users["Alex"]

	if correo, ok := users["Alex"]; ok {
		fmt.Println(correo)
	} else {
		fmt.Println("No fue posible obtener el valor ")
	}

}
