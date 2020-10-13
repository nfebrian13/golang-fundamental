package main

import (
	"fmt"
	"math"
)

/*
	fungsi adalah sekumpulan blok kode yang dibungkus dengan nama tertentu.
*/

func multiplereturn() {

	var diameter float64 = 15

	/* untuk menampung return 2 value maka variabel penampungnya juga harus 2 */
	var area, circumference = calculate(diameter)

	fmt.Printf("luas lingkaran\t\t: %.2f \n", area)
	fmt.Printf("keliling lingkaran\t: %.2f \n", circumference)

}

/* contoh fungsi untuk mengembalikan 2 value */
func calculate(d float64) (float64, float64) {
	// hitung luas
	var area = math.Pi * math.Pow(d/2, 2)
	// hitung keliling
	var circumference = math.Pi * d

	// kembalikan 2 nilai
	return area, circumference
}
