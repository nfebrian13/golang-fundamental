package main

import "fmt"

/*
	Konstanta adalah jenis variabel yang nilainya tidak bisa diubah.
	Inisialisasi nilai hanya dilakukan sekali di awal, setelah itu variabel tidak bisa diubah nilainya.
*/

func konstanta() {
	const lastName = "febriana"
	fmt.Print("halo ", lastName, "!\n")
}
