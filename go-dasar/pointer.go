package main

import (
	"fmt"
)

/*
	Pointer adalah reference atau alamat memori.
	Variabel pointer berarti variabel yang berisi alamat memori suatu nilai.
	Sebagai contoh sebuah variabel bertipe integer memiliki nilai 4,
	maka yang dimaksud pointer adalah alamat memori dimana nilai 4 disimpan, bukan nilai 4 itu sendiri.

	Variabel bertipe pointer ditandai dengan adanya tanda asterisk (*)
	tepat sebelum penulisan tipe data ketika deklarasi.
*/
/*
	Ada dua hal penting yang perlu diketahui mengenai pointer:

	Variabel biasa bisa diambil nilai pointernya, caranya dengan menambahkan tanda ampersand (&)
	tepat sebelum nama variabel. Metode ini disebut dengan referencing.
	Dan sebaliknya, nilai asli variabel pointer juga bisa diambil,
	dengan cara menambahkan tanda asterisk (*) tepat sebelum nama variabel.
	Metode ini disebut dengan dereferencing.*/

func pointer() {
	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println("numberA (value)   :", numberA)  // 4
	fmt.Println("numberA (address) :", &numberA) // 0xc20800a220

	fmt.Println("numberB (value)   :", *numberB) // 4
	fmt.Println("numberB (address) :", numberB)  // 0xc20800a220
}
