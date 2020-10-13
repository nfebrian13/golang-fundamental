package main

import "fmt"

func initialisasimake() {

	/* contoh inisialisasi array dengan make */

	var fruits = make([]string, 2)

	fruits[0] = "apple"
	fruits[1] = "mango"

	fmt.Println("index ke ", fruits)

}
