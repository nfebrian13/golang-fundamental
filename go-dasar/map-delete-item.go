package main

import "fmt"

/*
	delete() digunakan untuk menghapus item dengan key tertentu pada variabel map.
*/

func mapdeleteitem() {

	var chicken = map[string]int{
		"dada": 10,
		"paha": 8,
	}

	fmt.Println(len(chicken))
	fmt.Println(chicken)

	delete(chicken, "paha")

	fmt.Println(len(chicken))
	fmt.Println(chicken)
}
