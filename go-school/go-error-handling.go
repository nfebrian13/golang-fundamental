package handling

import (
	"errors"
	"fmt"
	"os"
)

/* error handling.*/

func handling() {
	// hourday := time.Now().Hour() /* type inference */
	greeting, err := getGreeting(6)

	if err != nil {
		os.Exit(1)
	}
	fmt.Println(greeting)
}

func getGreeting(hour int) (string, error) {
	var message string
	if hour < 7 {
		err := errors.New("Too early for greetings!")
		return message, err
	}
	if hour < 12 {
		message = "Good Morning"
	} else if hour < 18 {
		message = "Good Afternoon"
	} else {
		message = "Good Evening"
	}
	return message, nil
}
