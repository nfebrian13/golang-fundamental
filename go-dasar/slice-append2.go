package main

import "fmt"

/*
	hal hal yg perlu diketahui dalam penggunaan fungsi ini.
	1. ketika jumlah elemen dan lebar kapasitas adalah sama len(fruits) == cap(fruits)
	   maka elemen baru hasil append() merupakan referensi baru
	2. ketika jumlah elemen lebih kecil dibanding kapasitas len(fruits) < cap(fruits)
	   maka elemen baru akan ditempatkan kedalam cakupan kapasitas, menjadikan semua elemen
	   slice lain referensinya sama akan berubah nilanya.
*/

func append2() {

	var fruits = []string{"apple", "grape", "banana"}
	var bFruits = fruits[0:2]

	fmt.Println(cap(bFruits))
	fmt.Println(len(bFruits))

	fmt.Println()
	fmt.Println(fruits)
	fmt.Println(bFruits)

	fmt.Println()
	var cFruits = append(bFruits, "papaya")

	fmt.Println(fruits)
	fmt.Println(bFruits)
	fmt.Println(cFruits)
}
