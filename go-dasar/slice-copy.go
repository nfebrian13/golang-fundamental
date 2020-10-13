package main

import "fmt"

/*
	copy(dst,src) fungsi ini digunakan untuk men-copy elemen slice pada src ke dst parameter pertama.
*/

func slicecopy() {

	/* contoh pertama */
	dst := make([]string, 3)
	src := []string{"watermelon", "apple", "grape", "banana"}
	n := copy(dst, src)

	fmt.Println(dst)
	fmt.Println(src)
	fmt.Println(n)

	/* contoh kedua */
	dst1 := []string{"potato", "potato", "potato"}
	src1 := []string{"watermelon", "pinnaple"}
	n1 := copy(dst1, src1)

	fmt.Println(dst1)
	fmt.Println(src1)
	fmt.Println(n1)
}
