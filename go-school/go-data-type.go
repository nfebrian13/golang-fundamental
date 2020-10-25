package datatype

import (
	"fmt"
)

/* 2 cara untuk mendeklarasikan variable
1. type inference, tidak mendeklarasikan data type nya.
2. manual type declaration.*/

func datatype() {
	message1 := "Hello World from Nana!" /* type inference */
	fmt.Println(message1)

	var message2 = "Hello World from Tari!" /* manual type declaration */
	fmt.Println(message2)
}
