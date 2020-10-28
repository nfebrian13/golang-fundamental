package passbyreference

import "fmt"

func passbyreference() {
	message := "Hello World."
	sayHello(&message)
	fmt.Println(message)
}

func sayHello(message *string) {
	fmt.Println(*message)
	*message = "Hello Nana."
}
