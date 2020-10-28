package function

import "fmt"

func function() {

	/* example single return*/
	arr := []int{1, 3, 5, 9}
	hsl := add(arr)
	fmt.Println(hsl)

	/* example multiple return*/
	arr2 := []int{1, 3, 5, 9}
	term, hsl2 := added2(arr2)
	fmt.Println("term is ", term, " result is ", hsl2)
}

/* single return value */
func add(arr []int) int {
	result := 0
	for _, e := range arr {
		result += e
	}
	return result
}

/* multiple return value */
func added(arr []int) (int, int) {
	result := 0
	for _, e := range arr {
		result += e
	}
	return len(arr), result
}

/* multiple return value */
func added2(arr []int) (term int, sum int) {
	for _, e := range arr {
		sum += e
	}
	term = len(arr)
	return
}
