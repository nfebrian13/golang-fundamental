package main

import "fmt"

/*
	zero value dari map adalah nil, maka tiap variabel bertipe map harus di inisialisasi secara explisit
	nilai awalnya
*/

func mapinitialisasi() {

	/* initialisasi map berikut menyebabkan nil
	var data map[string]int
	data["one"] = 1 */

	/* initialisasi map berikut menyebabkan nil */
	var data map[string]int
	data = map[string]int{}
	data["one"] = 1

	fmt.Println("data one", data["one"])

	/* cara penulisan initialisasi map */
	//cara vertikal
	var chicken1 = map[string]int{"dada": 10, "paha": 8}
	fmt.Println("dada ", chicken1["dada"])

	//cara horizontal
	var chicken2 = map[string]int{
		"dada": 10,
		"paha": 8,
	}
	fmt.Println("dada ", chicken2["dada"])
}
