package main

import "fmt"

func switchcase() {
	var value int = 9

	switch value {
	case 9:
		fmt.Println("perfect")
	case 8:
		fmt.Println("awesome")
	default:
		fmt.Println("try again")
	}
}
