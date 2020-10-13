package main

import "fmt"

/* slice adalah reference elemen array.
   slice dapat dibuat ataupun dihasilkan dari manipulasi sebuah array ataupun slice lainya.
*/

func slice() {

	/* inisialisasi slice.
	   perbedaan slice dengan array pada letak inisialisasinya, pada deklarasi variabel
	   tidak dituliskan jumlah elemenya.
	*/
	var fruits = []string{"apple", "grape", "banana", "melon"}
	var newfruit = fruits[0:2]

	fmt.Println("index ke ", fruits[0])
	fmt.Println("semua elemen mulai indeks 0 dan sebelum indeks 2", newfruit)

	fmt.Println("semua elemen ", fruits[:])
	fmt.Println("semua elemen mulai indeks ke 2 ", fruits[2:])
	fmt.Println("semua elemen hingga sebelum indeks ke 2", fruits[:2])
}
