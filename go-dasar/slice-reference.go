package main

import "fmt"

/* slice adalah merupakan tipe data reference artinya jika ada slice baru yg terbentuk dari slice lama
   maka data elemen slice yg baru akan memiliki alamat memori yg sama dengan elemen slice yg lama.
*/

func slicereference() {

	var fruits = []string{"apple", "grape", "banana", "melon"}
	var aFruits = fruits[0:3]
	var bFruits = fruits[1:4]

	var aaFruits = aFruits[1:2]
	var bbFruits = aFruits[0:1]

	fmt.Println(fruits)
	fmt.Println(aFruits)
	fmt.Println(bFruits)
	fmt.Println(aaFruits)
	fmt.Println(bbFruits)

	fmt.Println()

	// buah "grape" di ubah menjadi "pinnaple"
	bbFruits[0] = "pinnaple"

	fmt.Println(fruits)
	fmt.Println(aFruits)
	fmt.Println(bFruits)
	fmt.Println(aaFruits)
	fmt.Println(bbFruits)
}
