package loop

import (
	"fmt"
)

func loop() {
	faktor := factorial(3)
	fmt.Println(faktor)

	fmt.Println()
	printValue()
}

func factorial(n int) int {
	var hsl int = 1
	for i := 1; i <= n; i++ {
		hsl *= i
	}
	return hsl
}

/* another example loop using for */
func printValue() {
	i := 0
	checkvalue := true
	for checkvalue {
		if i > 5 {
			checkvalue = false
		}
		fmt.Println(i)
		i++
	}
}
