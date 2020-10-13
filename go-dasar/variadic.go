package main

import (
	"fmt"
	"strings"
)

/*
	variadic adalah pembuatan fungsi dengan parameter sejenis yang tidak terbatas.
	tidak terbatas maksudnya jumlah dari parameternya dapat berapa saja.
*/

func variadic() {

	var avg = cal(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
	var msg = fmt.Sprintf("Rata-rata : %.2f", avg)
	fmt.Println(msg)

	/* contoh initialisasi slice yang digunakan untuk parameter variadic */
	var numbers = []int{2, 4, 3, 5, 4, 3, 3, 5, 5, 3}
	var rata = cal(numbers...)
	var msg1 = fmt.Sprintf("Rata-rata : %.2f", rata)

	fmt.Println(msg1)

	var hobbies = []string{"sleeping", "eating"}
	yourHobbies("wick", hobbies...)
}

/* fungsi dengan parameter biasa dan variadic */

func yourHobbies(name string, hobbies ...string) {
	var hobbiesAsString = strings.Join(hobbies, ", ")

	fmt.Printf("Hello, my name is: %s\n", name)
	fmt.Printf("My hobbies are: %s\n", hobbiesAsString)
}

func cal(numbers ...int) float64 {
	var total int = 0
	for _, number := range numbers {
		total += number
	}

	var avg = float64(total) / float64(len(numbers))
	return avg
}
