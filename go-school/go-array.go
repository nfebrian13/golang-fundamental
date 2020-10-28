package main

import (
	"fmt"
)

func main() {
	langs := proglanguage()

	// for i := range langs {
	// 	fmt.Println(langs[i])
	// }

	for _, element := range langs {
		fmt.Println(element)
	}
}

func proglanguage() [4]string {
	lang := [4]string{"Go", "Java", "Ruby", "Javascript"}
	return lang
}
