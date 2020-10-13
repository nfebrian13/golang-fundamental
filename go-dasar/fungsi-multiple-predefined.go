package main

import (
	"fmt"
	"math"
)

/*
	fungsi predefined adalah return value yang didefiniskan di awal fungsi
*/

func predefined() {

	var diameter float64 = 15

	/* untuk menampung return 2 value maka variabel penampungnya juga harus 2 */
	var area, circumference = calculate2(diameter)

	fmt.Printf("luas lingkaran\t\t: %.2f \n", area)
	fmt.Printf("keliling lingkaran\t: %.2f \n", circumference)
}

/* contoh fungsi untuk mengembalikan 2 value */
func calculate2(d float64) (area float64, circumference float64) {
	// hitung luas
	area = math.Pi * math.Pow(d/2, 2) // math.Pow digunakan untuk memangkat nilai
	// hitung keliling
	circumference = math.Pi * d // math.Pi nilai constant dari 22/7

	// kembalikan 2 nilai
	return
}
