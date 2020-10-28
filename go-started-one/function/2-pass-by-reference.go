package main

import "fmt"

func main() {
	message := "Hello World."
	sayHello(&message)
	fmt.Println(message)
}

func sayHello(message *string) {
	fmt.Println(*message)
	*message = "Hello Nana."
}
