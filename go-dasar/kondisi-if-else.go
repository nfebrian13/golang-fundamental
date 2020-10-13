package main

import "fmt"

func ifelse() {
	var value int = 8

	if value >= 8 {
		fmt.Println("excelent")
	} else if value > 6 && value < 8 {
		fmt.Println("not bad")
	} else {
		fmt.Println("try again")
	}
}
