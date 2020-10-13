package main

import "fmt"

func underscore() {

	/* array dengan dengan underscore (_) digunanakan untuk
	   apabila didalam array contoh dibawah hanya
	   akan menampilkan valuenya saja tidak dengan indeknya
	   maka menggunakan array underscore _ */

	var fruits = [4]string{"apel", "jambu", "alpukat", "melon"}

	for _, fruit := range fruits {
		fmt.Println("index ke ", fruit)
	}

}
