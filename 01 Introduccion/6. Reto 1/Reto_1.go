package main

import "fmt"

func practica1(num1, num2 int) int {
	return num1 + num2
}
func practica2_1(num1, num2 int) int {
	return num1 / num2
}
func practica2_2(num1, num2 int) int {
	return num1 % num2
}
func practica3(num1 int) float64 {
	const porcentaje float64 = 0.18
	var igv float64 = float64(num1) * porcentaje
	return igv
}

func main() {
	var numero1 int
	var numero2 int
	fmt.Print("Seleccionar practica (1, 2, 3): ")
	var seleccion int
	fmt.Scanln(&seleccion)

	if (seleccion == 1 || seleccion == 2) {
		fmt.Print("Ingrese el primer numero: ")
		fmt.Scanln(&numero1)

		fmt.Print("Ingrese el segundo numero: ")
		fmt.Scanln(&numero2)

		if (seleccion == 1){
			suma := practica1(numero1, numero2)
			fmt.Printf("La suma es: %v", suma)
		} else {
			cociente := practica2_1(numero1, numero2)
			residuo := practica2_2(numero1, numero2)

			fmt.Printf("El ciciente es: %v \n", cociente)
			fmt.Printf("El residuo es: %v", residuo)
		}
	} else {
		fmt.Print("Ingrese el valor de venta: ")
		fmt.Scanln(&numero1)

		IGV := practica3(numero1)
		venta := float64(numero1) + IGV

		fmt.Printf("El IGV: %v \n", IGV)
		fmt.Printf("El precio de venta: %v \n", venta)

	}


}