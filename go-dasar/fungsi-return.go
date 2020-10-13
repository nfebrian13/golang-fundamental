package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	fungsi adalah sekumpulan blok kode yang dibungkus dengan nama tertentu.
*/

func fungsireturn() {

	/* fungsi ini digunakan untuk memastikan bahwa angka random yang akan digenerate benar-benar acak */
	rand.Seed(time.Now().Unix())
	var randomValue int

	randomValue = randomWithRange(2, 10)
	fmt.Println("random number is ", randomValue)

	fmt.Println()

	divideNumber(10, 2)
	divideNumber(8, 0)
	divideNumber(8, -4)

}

func randomWithRange(min, max int) int {
	var value = rand.Int()%(max-min+1) + min
	return value
}

/* contoh return untuk menghentikan proses dalam blok fungsi */
func divideNumber(m, n int) {
	if n == 0 {
		fmt.Printf("invalid divider. %d cannot divided by %d\n", m, n)
		return
	}

	var res = m / n
	fmt.Printf("%d / %d = %d\n", m, n, res)

}
