package main

import "fmt"

/*
	delete() digunakan untuk menghapus item dengan key tertentu pada variabel map.
*/

func mapexist() {

	var chicken = map[string]int{
		"dada": 10,
		"paha": 8,
	}

	var value, isExist = chicken["paha"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exist")
	}

}
