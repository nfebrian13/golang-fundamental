package main

import (
	"fmt"
)

/*
Struct adalah kumpulan definisi variabel (atau property) dan atau fungsi (atau method),
yang dibungkus sebagai tipe data baru dengan nama tertentu.
Property dalam struct, tipe datanya bisa bervariasi.
Mirip seperti map, hanya saja key-nya sudah didefinisikan di awal,
dan tipe data tiap itemnya bisa berbeda
*/

func structexample() {
	var s1 hewan
	s1.name = "Kucing"
	s1.umur = 2

	fmt.Println("name  :", s1.name)
	fmt.Println("grade :", s1.umur)
}

type hewan struct {
	name string
	umur int
}
