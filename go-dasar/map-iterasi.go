package main

import "fmt"

/*
	iterasi item map menggunakan for - range
*/

func iterasimap() {

	var chicken = map[string]int{
		"dada":  10,
		"paha":  8,
		"sayap": 6,
		"ceker": 5,
	}

	for key, val := range chicken {
		fmt.Println(key, " \t:", val)
	}

}
