package main

import "fmt"

func main() {

	numero := 4

	if numero == 45 {
		fmt.Println("Aplicando IF")
	} else if numero == 40 {
		fmt.Println("Else IF ejecutado")
	} else {
		fmt.Println("Ejecutando ELSE")
	}

	edad := 30

	if edad >= 18 {
		fmt.Println("Eres mayor de edad")
	} else if edad < 18 && edad >= 0 {
		fmt.Println("Eres menor de edad")
	} else {
		fmt.Println("La edad no es valida")
	}
}
