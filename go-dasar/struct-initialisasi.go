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

func initialisasistruct() {

	/* cara pertama */
	var s1 = student{}
	s1.name = "wick"
	s1.grade = 2

	/* cara kedua */
	var s2 = student{"ethan", 2}

	/* cara ketiga */
	var s3 = student{name: "jason"}

	fmt.Println("student 1 :", s1.name)
	fmt.Println("student 2 :", s2.name)
	fmt.Println("student 3 :", s3.name)
}

type student struct {
	name  string
	grade int
}
