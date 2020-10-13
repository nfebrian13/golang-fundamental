package main

import "fmt"

/* cap() fungsi berikut digunakan untuk menghitung lebar atau kapasitas maksimum slice */

func slicecap() {

	var fruits = []string{"apple", "grape", "banana", "melon"}
	fmt.Println("index len is: ", len(fruits))
	fmt.Println("index cap is: ", cap(fruits))

	var aFruits = fruits[0:3]
	fmt.Println("index len is: ", len(aFruits))
	fmt.Println("index cap is: ", cap(aFruits))

	var bFruits = fruits[1:4]
	fmt.Println("index len is: ", len(bFruits))
	fmt.Println("index cap is: ", cap(bFruits))
}
