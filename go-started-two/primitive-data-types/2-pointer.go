package pointer

import "fmt"

func pointer() {
	/* tanda * adalah pointer */
	var firstName *string = new(string)
	*firstName = "Nana"
	fmt.Println(firstName)  // print memory address
	fmt.Println(*firstName) // print value

	lastName := "febriana"
	fmt.Println(lastName)

	ptr := &lastName
	fmt.Println(ptr, *ptr)

	lastName = "retno"
	fmt.Println(ptr, *ptr)
}
