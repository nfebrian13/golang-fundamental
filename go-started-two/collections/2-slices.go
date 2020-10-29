package slices

import "fmt"

func slices() {
	arr := [3]int{1, 2, 3}
	slice := arr[:]

	arr[1] = 42
	slice[2] = 27

	fmt.Println()
	fmt.Println(arr, slice)

	expslice := []int{1, 2, 3}
	fmt.Println(expslice)

	expslice = append(expslice, 4, 42, 27)
	fmt.Println(expslice)

	s2 := expslice[1:]
	s3 := expslice[:2]
	s4 := expslice[1:2]

	fmt.Println()
	fmt.Println("s2 is: ", s2, " s3 is: ", s3, " s4 is: ", s4)
}
