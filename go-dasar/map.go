package main

import "fmt"

/*
	map adalah tipe data asosiatif yang ada di go, berbentuk key-value pair.
*/

func mapexp() {

	var ayam map[string]int
	ayam = map[string]int{}

	ayam["dada"] = 10
	ayam["paha"] = 8

	fmt.Println("dada", ayam["dada"])
	fmt.Println("paha", ayam["paha"])
}
