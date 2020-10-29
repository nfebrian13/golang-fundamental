package iotaexample

import "fmt"

const (
	first = iota
	second
)

const (
	third = iota
)

func iotaexample() {
	fmt.Println(first, second, third)
}
