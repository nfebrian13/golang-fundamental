package main

import (
	"fmt"
)

/*
	closure adalah sebuah fungsi yang bisa disimpan dalam variabel.
	Dengan menerapkan konsep tersebut, kita bisa membuat fungsi didalam fungsi,
	atau bahkan membuat fungsi yang mengembalikan fungsi.
*/

func closure() {

	var getMinMax = func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max
	}

	var numbers = []int{2, 3, 4, 3, 4, 2, 3}
	var min, max = getMinMax(numbers)
	fmt.Printf("data : %v\nmin  : %v\nmax  : %v\n", numbers, min, max) // %v digunakan untuk menampilkan segala jenis data
}
