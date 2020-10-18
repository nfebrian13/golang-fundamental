package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

type student struct {
	grade  int
	age    int // embeed struct dengan properti yg sama
	person     // Embedded struct
}

func substruct() {
	var p1 = person{name: "wick", age: 21}
	var s1 = student{person: p1, grade: 2}

	fmt.Println("name  :", s1.name)
	fmt.Println("age   :", s1.age)
	fmt.Println("grade :", s1.grade)
}