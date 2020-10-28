package collection

import (
	"fmt"
)

func collection() {
	arr := [3]int{}
	arr[0] = 13
	arr[1] = 2
	arr[2] = 93

	/* slice example */
	myArr := [...]int{13, 2, 93}
	mySlice := myArr[:]
	fmt.Println(myArr)

	/* berikut adalah contoh menambahkan element pada array mySlice */
	mySlice = append(mySlice, 100)
	fmt.Println(myArr)
	fmt.Println(mySlice)
}
