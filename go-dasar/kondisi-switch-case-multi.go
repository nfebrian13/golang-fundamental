package main

import "fmt"

func casemulti() {
	var value int = 10

	switch value {
	case 9:
		fmt.Println("perfect")
	case 8, 7, 6, 5:
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you can be better!")
		}
	}
}
