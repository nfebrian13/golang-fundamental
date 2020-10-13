package main

import "fmt"

func mainawal() {

	/* array */

	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	fmt.Println("Jumlah element \t\t", len(fruits))
	fmt.Println("Isi semua element \t", fruits)

	fmt.Println("Print index ke 0 \t", fruits[0])
	fmt.Println("Print index ke 1 \t", fruits[1])

}
