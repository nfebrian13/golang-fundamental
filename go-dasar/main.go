package contoh

import "fmt"

func contoh() {

	var s1 = student{age:30, name : "dwi"}
	var s student
	s.age = 20
	s.name = "Nana"

	s.sayHello()
	s1.sayHello()
}

type student struct {
	name string
	age  int
}

func (s student) sayHello() {
	fmt.Println("Hello ", s.name)
}
