package main

import "fmt"

/*
	== apakah nilai kiri sama dengan nilai kanan
	!= apakah nilai kiri tidak sama dengan nilai kanan
	<  apakah nilai kiri lebih kecil daripada nilai kanan
	<= apakah nilai kiri lebih kecil atau sama dengan nilai kanan
	>  apakah nilai kiri lebih besar dari nilai kanan
	>= apakah nilai kiri lebih besar atau sama dengan nilai kanan
*/

func perbandingan() {
	var value = (((2+6)%3)*4 - 2) / 3
	var isEqual = (value == 2)

	fmt.Printf("nilai %d (%t) \n", value, isEqual)
}
