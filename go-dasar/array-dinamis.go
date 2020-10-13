package main

import "fmt"

func arraydinamis() {

	/* inisialisasi nilai awal array tanpa jumlah elemen */

	var numbers = [...]int{1, 2, 3, 4, 5}

	fmt.Println("Data array \t\t", numbers)
	fmt.Println("Jumlah element \t\t", len(numbers))

}
