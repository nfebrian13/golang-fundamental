package main

import "fmt"

/*
	Keyword fallthrough digunakan untuk memaksa proses pengecekkan diteruskan ke case selanjutnya
	dengan tanpa menghiraukan nilai kondisinya, jadi case di pengecekan selanjutnya tersebut
	selalu dianggap benar (meskipun aslinya adalah salah).
*/

func main() {
	var point int = 6

	switch {

	case point == 8:
		fmt.Println("perfect")
	case (point < 8) && (point > 3):
		fmt.Println("awesome")
		fallthrough
	case point < 5:
		fmt.Println("you need to learn more")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}
}
