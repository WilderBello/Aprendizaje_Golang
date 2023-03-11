package main

import "fmt"

func main() {
	frutas := []string{"Pera", "Manzana", "Naranja"}

	frutas = append(frutas, "Melocoton", "Melon")

	frutas[0] = "Sandia"

	for i := 0; i < len(frutas); i++ {
		fmt.Println(frutas[i])
	}

	for i := 0; i < len(frutas); i++ {
		if frutas[i] == "Melocoton" {
			fmt.Println("Hay una coincidencia", i)
		}
	}
}
