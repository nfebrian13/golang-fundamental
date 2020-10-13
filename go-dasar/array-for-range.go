package main

import "fmt"

func forrange() {

	/* array dengan perulangan */

	var fruits = [4]string{"apel", "jambu", "alpukat", "melon"}

	for i, fruit := range fruits {
		fmt.Println("index ke ", i, "\t\t", fruit)
	}

}
