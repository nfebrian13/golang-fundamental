package anonymous

import "fmt"

func anonymous() {

	/* example anonymous function */
	added := func(arr ...int) (term int, sum int) {
		for _, e := range arr {
			sum += e
		}
		term = len(arr)
		return
	}
	term, sum := added(1, 2)
	fmt.Println("term is ", term, " result is ", sum)
}
