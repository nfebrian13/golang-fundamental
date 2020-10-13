package main

import "fmt"

func perulangan() {

	/* 1. penulisan for dengan cara pertama */
	for i := 0; i < 5; i++ {
		fmt.Println("1. penulisan for dengan cara pertama!")
	}

	/* 2. penulisan for dengan cara kedua (mirip while) */
	var i = 0

	for i < 5 {
		fmt.Println("2. penulisan for dengan cara kedua!")
		i++
	}

	/* 3. penulisan for dengan cara ketiga, tanpa argumen*/
	var j = 0
	for {
		fmt.Println("3. penulisan for dengan cara ketiga!")
		j++

		if j == 5 {
			break
		}
	}

	/* 4. penulisan for dengan cara keempat, break dan continue*/
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}
		if i > 8 {
			break
		}
		fmt.Println("angka", i)
	}

	/* 5. penulisan for dengan cara kelima, perulangan bersarang*/
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}

}
