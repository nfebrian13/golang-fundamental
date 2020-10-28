package constant

import (
	"fmt"
)

const (
	first = iota
	second
	third
)

func constant() {
	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(third)
}
