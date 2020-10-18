package main

import (
	"fmt"
	"strings"
)

/*
	closure adalah sebuah fungsi yang bisa disimpan dalam variabel.
	Dengan menerapkan konsep tersebut, kita bisa membuat fungsi didalam fungsi,
	atau bahkan membuat fungsi yang mengembalikan fungsi.
*/

func fungsisebagaiparameter() {
	var data = []string{"wick", "jason", "ethan"}
	var dataContainsO = filter(data, func(each string) bool {
		return strings.Contains(each, "o")
	})
	var dataLenght5 = filter(data, func(each string) bool {
		return len(each) == 5
	})

	fmt.Println("data asli \t\t:", data)
	// data asli : [wick jason ethan]

	fmt.Println("filter ada huruf \"o\"\t:", dataContainsO)
	// filter ada huruf "o" : [jason]

	fmt.Println("filter jumlah huruf \"5\"\t:", dataLenght5)
	// filter jumlah huruf "5" : [jason ethan]
}

/* Parameter callback merupakan sebuah closure yang dideklarasikan bertipe func(string) bool.
   Closure tersebut dipanggil di tiap perulangan dalam fungsi filter().
*/

func filter(data []string, callback func(string) bool) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}
