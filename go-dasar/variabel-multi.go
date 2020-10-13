package main

import (
	"fmt"
)

/*
	Go mengadopsi mendukung deklarasi banyak variabel
*/

func multi() {

	/*
		Berikut adalah contoh penulisan deklarsi banyak variabel dg 2 jenis penulisan
	*/
	var one, two, three, four string = "one", "two", "three", "four"
	var five, six string
	five = "5"
	six = "6"

	fmt.Println(one, two, three, four)
	fmt.Println(five, six)
}
