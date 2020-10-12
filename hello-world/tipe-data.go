package main

import (
	"fmt"
)

/*
	1. uint, tipe data untuk bilangan cacah (bilangan positif).
	2. int, tipe data untuk bilangan bulat (bilangan negatif dan positif).
	3. bool, ada 2 nilai yaitu true atau false.
	4. string, ciri adanya tanda petik "".
	5. nil bukan merupakan tipe data, melainkan sebuah nilai.
	   Variabel yang isi nilainya nil berarti memiliki nilai kosong.
	   ada beberapa tipe data yg dapat menggunakan nilai nil.
	   pointer, tipe data fungsi, slice, map, channel, interface kosong
*/

func main() {
	var positiveNumber uint8 = 89
	var negativeNumber int = -1243423644

	fmt.Println("positiveNumber ", positiveNumber)
	fmt.Println("negativeNumber ", negativeNumber)

	var exist bool = true
	fmt.Printf("exist? %t \n", exist)

	var message string = "Hallo"
	fmt.Printf("message: %s \n", message)
}
