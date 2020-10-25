package main

import (
	"fmt"
	"time"
)

/* function.*/

func main() {
	hourday := time.Now().Hour() /* type inference */
	greeting := getGreeting(hourday)
	fmt.Println(greeting)
}

func getGreeting(hour int) string {
	if hour < 12 {
		return "Good Morning"
	} else if hour < 18 {
		return "Good Afternoon"
	} else {
		return "Good Evening"
	}
}
