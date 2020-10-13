package main

import "fmt"

func gaya() {

	var fruits [4]string

	/* array dengan gaya horizontal
	fmt.Println("array dengan gaya horizontal")

	fruits = [4]string{"apple", "grape", "banana", "melon"}
	fmt.Println("Jumlah element \t\t", len(fruits)) */

	/* array dengan gaya vertical */
	fmt.Println("array dengan gaya horizontal")
	fruits = [4]string{
		"apple",
		"grape",
		"banana",
		"melon",
	}

	fmt.Println("Jumlah element \t\t", len(fruits))
}
