package main

import (
	"go-webservice/controllers"
	"net/http"
)

func main() {

	/*	fmt.Println("Hello, Gopher!")
		u := models.User{
			ID:        2,
			FirstName: "Nana",
			LastName:  "Febriana",
		}

		fmt.Println(u) */

	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
