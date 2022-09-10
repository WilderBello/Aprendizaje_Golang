package main

import "fmt"

func main() {
	//Slicen, son mutables, por lo tanto se puede cambiar la longitud
	numeros := []int{1, 2, 3, 4, 5}
	fmt.Println(numeros, len(numeros))

	//Agregar elementos con append
	numeros = append(numeros, 6) //Se guarda de nuevo en el mismo slicen
	numeros = append(numeros, 7) 
	fmt.Println(numeros, len(numeros))

	//Sub Slicen
	subSlicen := numeros[:2] //Se modifica al modificar el Slicen principal

	numeros[0] = 100

	fmt.Println(subSlicen)
	fmt.Println(numeros)

	//Punteros: Saber donde inicia y donde termina
	//Longitud
	//Capacidad

	meses := []string{"Enero", "Febrero", "Marzo"}

	fmt.Printf("Len: %v, Cap: %v, %p \n", len(meses), cap(meses), meses)

	meses = append(meses, "Abril")
	fmt.Printf("Len: %v, Cap: %v, %p \n", len(meses), cap(meses), meses) // La capacidad se duplica, cambia la referencia de memoria por ser otro slicen
}