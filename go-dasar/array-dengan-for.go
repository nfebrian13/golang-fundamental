package main

import "fmt"

func arrayfor() {

	/* array dengan perulangan */

	var fruits = [4]string{"apel", "jambu", "alpukat", "melon"}

	for i := 0; i < len(fruits); i++ {
		fmt.Println("index ke ", i, "\t\t", fruits[i])
	}

}
