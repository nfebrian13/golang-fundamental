package main

import "fmt"

/* cap() fungsi berikut digunakan untuk menambahkan elemen pada slice.
   elemen tersebut diposisikan setelah indeks paling akhir. */

func sliceappend() {

	var fruits = []string{"apple", "grape", "banana", "melon"}
	var cFruits = append(fruits, "papaya")

	fmt.Println(fruits)
	fmt.Println(cFruits)
}
