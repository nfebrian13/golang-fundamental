package main

import (
	"fmt"
)

func main() {
	arr := [3] int {}
	arr[0] = 13
	arr[1] = 2
	arr[2] = 93
	fmt.Println(arr)

	for _, element := range arr {
		fmt.Println(element)
	}
}
