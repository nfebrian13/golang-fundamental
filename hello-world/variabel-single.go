package main

import (
	"fmt"
)

/*
	Go mengadopsi 2 cara penulisan variabel
	1. menuliskan tipe data
	2. tidak menuliskan tipe data
*/

func variabel() {

	var firstName string = "Nana" // 1. menuliskan tipe data
	lastName := "Febriana"        // 2. tidak menuliskan tipe data (tanda = harus ditambah :)
	lastName = "Tari"             // asignment nilai selanjutnya tidak perlu menambahkan :

	fmt.Printf("halo %s %s!\n", firstName, lastName)

}
