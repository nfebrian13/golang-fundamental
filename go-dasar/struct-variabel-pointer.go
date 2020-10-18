package main

import (
	"fmt"
)

func structpointer() {
	var s1 = empoloyee{name: "wick", age: 2}

	var s2 *empoloyee = &s1
	fmt.Println("empoloyee 1, name :", s1.name)
	fmt.Println("empoloyee 4, name :", s2.name)

	s2.name = "ethan"
	fmt.Println("empoloyee 1, name :", s1.name)
	fmt.Println("empoloyee 4, name :", s2.name)
}

type empoloyee struct {
	name string
	age  int
}
