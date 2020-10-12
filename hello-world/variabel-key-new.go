package main

import (
	"fmt"
)

/*
	'new' digunakan untuk membuat variabel 'pointer' dg tipe data tertentu
	 nilai data defaultnya akan menyesuaikan tipe datanya
*/

func keynew() {

	name := new(string)

	fmt.Println(name)
	fmt.Println(*name) // '*' pendanda menggunakan pointer
}
