package struct

import "fmt"

func struct() {

	type user struct {
		id        int
		firstName string
		lastName  string
	}

	var u user
	u.id = 1
	u.firstName = "Nana"
	u.lastName = "Febriana"
	fmt.Println(u)

	u2 := user{id: 1, firstName: "Nana", lastName: "Febriana"}
	fmt.Println(u2)
}
