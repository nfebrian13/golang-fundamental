package arithmetic

import (
	"fmt"
)

func arithmetic() {

	add := 1 + 2
	fmt.Println(add)

	substract := 10 - 5
	fmt.Println(substract)

	modulo := 5 % 2
	fmt.Println(modulo)

	inc := 1
	inc++
	fmt.Println(inc)
	inc++
	fmt.Println(inc)
	inc += 5
	fmt.Println(inc)
	inc *= 10
	fmt.Println(inc)
}
