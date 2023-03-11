package main

import "fmt"

func main() {
	suma := 10 + 20
	resta := 10 - 5
	multiplicacion := 10 * 5
	division := 10 / 5
	resto := 10 % 5

	fmt.Println(suma, resta, multiplicacion, division, resto)
	suma++ // suma + 1
	fmt.Println(suma)
	suma-- // suma - 1
	fmt.Println(suma)
	suma += 10 // suma = suma + 10
	fmt.Println(suma)

	fmt.Println(10 > 5)
	fmt.Println(10 < 5)
	fmt.Println(10 >= 10)
	fmt.Println(10 <= 5)
	fmt.Println(10 == 10)
	fmt.Println(10 != 5)
	fmt.Println(10 < 5 && 10 > 4)
	fmt.Println(10 > 5 || 10 < 4)
}
