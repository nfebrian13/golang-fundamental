package main

import "fmt"

func multidimensi() {

	/* inisialisasi nilai awal array multidimensi */

	var numbers1 = [2][3]int{[3]int{3, 3, 2}, [3]int{3, 4, 5}}

	var numbers2 = [2][3]int{{3, 3, 2}, {3, 4, 5}}

	fmt.Println("numbers1 \t\t", len(numbers1))
	fmt.Println("numbers1 \t\t", numbers1)

	fmt.Println("numbers1 \t\t", len(numbers2))
	fmt.Println("numbers1 \t\t", numbers1)
}
