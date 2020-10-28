package passbyvalue

import "fmt"

func passbyvalue() {
	message := "Hello Nana!"
	sayHello(message)
}

func sayHello(message string) {
	fmt.Println(message)
}
