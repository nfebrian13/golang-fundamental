package main

import "fmt"

/* len() fungsi berikut digunakan untuk menghitung jumlah elemen slice yg ada. */

func fungsilen() {

	var fruits = []string{"apple", "grape", "banana", "melon"}
	fmt.Println("index is: ", len(fruits))
}
