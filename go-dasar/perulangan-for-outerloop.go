package main

import "fmt"

func forouterloop() {

	/* outerLoop: adalah ketika kondisi i == 3 maka break dipanggil dengan target adalah
	   perulangan yang dilabeli outerLoop perulangan tersebut akan dihentikan
	*/

outerLoop:
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}
}
