package main

import "fmt"

func main() {

	/* App para restaurante
	Descuentos de 10% hasta 100 de consumo
	Descuento de 20% hasta 200 de consumo
	Descuento de 30% mayor 300 de consumo
	*/

	var Total float64
	var Descuento float64
	var Consumo float64

	fmt.Print("Ingrese el valor del consumo: ")
	fmt.Scanln(&Consumo)

	if Consumo > 0 {
		if Consumo <= 100 {
			Descuento = Consumo * 0.1
		} else if Consumo > 100 && Consumo <= 200 {
			Descuento = Consumo * 0.2
		} else {
			Descuento = Consumo * 0.3
		}

		Total = Consumo - Descuento

		fmt.Printf("El total a pagar es %v con un descuento de %v", Total, Descuento)

	} else {
		fmt.Print("Valor ingresado incorrectamente...")
	}

}
