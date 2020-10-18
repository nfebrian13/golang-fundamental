package main

import (
	"fmt"
)

/*
	Ketika salah satu variabel pointer di ubah nilainya,
	sedang ada variabel lain yang memiliki referensi memori yang sama,
	maka nilai variabel lain tersebut juga akan berubah
*/

func pointerefek() {
	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println("numberA (value)   :", numberA)
	fmt.Println("numberA (address) :", &numberA)
	fmt.Println("numberB (value)   :", *numberB)
	fmt.Println("numberB (address) :", numberB)

	fmt.Println("")

	numberA = 5

	fmt.Println("numberA (value)   :", numberA)
	fmt.Println("numberA (address) :", &numberA)
	fmt.Println("numberB (value)   :", *numberB)
	fmt.Println("numberB (address) :", numberB)
}

/*
	Variabel numberA dan numberB memiliki referensi memori yang sama.
	Perubahan pada salah satu nilai variabel tersebut akan memberikan efek pada variabel lainnya.
	Pada contoh di atas, numberA nilainya di ubah menjadi 5.
	membuat nilai asli variabel numberB ikut berubah menjadi 5.
*/
