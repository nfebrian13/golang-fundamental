package main

import "fmt"

/*
	+ penambahan
	- pengurangan
	* perkalian
	/ pembagian
	% modulus / sisa hasil pembagian
*/

func aritmetika() {
	var value = (((2+6)%3)*4 - 2) / 3
	fmt.Print("value is ", value, "!\n")
}
