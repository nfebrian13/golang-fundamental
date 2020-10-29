package main

import (
	"fmt"
	"go-webservice/models"
)

func main() {
	fmt.Println("Hello, Gopher!")
	u := models.User{
		ID:        2,
		FirstName: "Nana",
		LastName:  "Febriana",
	}

	fmt.Println(u)
}
