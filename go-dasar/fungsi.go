package main

import (
	"fmt"
	"strings"
)

/*
	fungsi adalah sekumpulan blok kode yang dibungkus dengan nama tertentu.
*/

func belajarfungsi() {
	var names = []string{"nana", "febriana"}
	printMessage("hallo", names)
}

func printMessage(message string, arr []string) {
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)
}
