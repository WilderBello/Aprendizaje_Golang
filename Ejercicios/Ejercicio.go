// Ejercicio que recorra los numeros del 0 al 20 y me diga los
// que son pares
package main

import "fmt"

func main() {
	var cantidad int = 20

	var numeros []int

	for i := 0; i <= cantidad; i++ {
		if i%2 == 0 {
			numeros = append(numeros, i)
			//fmt.Println(i)
		}

	}

	fmt.Println(numeros)
}
