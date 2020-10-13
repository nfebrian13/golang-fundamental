package main

import "fmt"

/*
	delete() digunakan untuk menghapus item dengan key tertentu pada variabel map.
*/

func kombinasislice() {

	/* */
	var chickens = []map[string]string{
		map[string]string{"name": "chicken blue", "gender": "male"},
		map[string]string{"name": "chicken red", "gender": "male"},
		map[string]string{"name": "chicken yellow", "gender": "female"},
	}

	for _, chicken := range chickens {
		fmt.Println(chicken["gender"], chicken["name"])
	}

	fmt.Println()

	/* deklarasi seperti diatas go versi baru */
	var ayam = []map[string]string{
		{"name": "chicken blue", "gender": "male"},
		{"name": "chicken red", "gender": "male"},
		{"name": "chicken yellow", "gender": "female"},
	}

	for _, chicken := range ayam {
		fmt.Println(chicken["gender"], chicken["name"])
	}

	fmt.Println()

	/* dalam map tiap elemen bisa saja memiliki key yang berbeda beda */
	var hewan = []map[string]string{
		{"name": "chicken blue", "gender": "male"},
		{"address": "jababeka street", "id": "k001"},
		{"community": "chicken lovers"},
	}

	for _, chicken := range hewan {
		fmt.Println(chicken["gender"], chicken["name"])
	}
}
