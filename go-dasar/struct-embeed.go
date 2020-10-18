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

func structembeed() {
	var s1 = student{}
	s1.name = "wick"
	s1.age = 21
	s1.person.age = 21
	s1.grade = 2

	fmt.Println("name  :", s1.name)
	fmt.Println("age   :", s1.age)
	fmt.Println("age person   :", s1.person.age)
	fmt.Println("grade :", s1.grade)
}
