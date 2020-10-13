package main

import "fmt"

/*
	3 indeks adalah teknik slicing elemen yg sekaligus menentukan kapasitasnya.
	angka kapasitas yang diisikan tidak boleh melebihi kapasitas slice yang akan
	di slicing.
*/

func slicetigaindeks() {

	var fruits = []string{"watermelon", "apple", "grape", "banana"}
	var aFruits = fruits[0:2]
	var bFruits = fruits[0:2:2]

	fmt.Println(fruits)
	fmt.Println(len(fruits))
	fmt.Println(cap(fruits))

	fmt.Println(aFruits)
	fmt.Println(len(aFruits))
	fmt.Println(cap(aFruits))

	fmt.Println(bFruits)
	fmt.Println(len(bFruits))
	fmt.Println(cap(bFruits))
}
