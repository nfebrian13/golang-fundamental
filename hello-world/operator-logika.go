package main

import (
	"fmt"
)

/*
	&& kiri dan kanan
	|| kiri atau kanan
	!  negasi / nilai kebalikan
*/
func main() {
	var left bool = false
	var right bool = true

	var leftAndRight = left && right
	fmt.Printf("left && right \t(%t) \n", leftAndRight)

	var leftOrRight = left || right
	fmt.Printf("left || right \t(%t) \n", leftOrRight)

	var leftReverse = !left
	fmt.Printf("!left \t\t(%t) \n", leftReverse)

}
