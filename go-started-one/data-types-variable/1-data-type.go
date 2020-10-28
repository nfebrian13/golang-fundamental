package datatype

import (
	"fmt"
)

func datatype() {
	var myInt int = 42
	fmt.Println(myInt)

	var myFloat32 float32 = 42.
	fmt.Println(myFloat32)
	println(myFloat32)

	myString := "Hello Go"
	fmt.Println(myString)

	myComplex := complex(2, 3)
	fmt.Println(myComplex)
	fmt.Println(real(myComplex))
	fmt.Println(imag(myComplex))
}
