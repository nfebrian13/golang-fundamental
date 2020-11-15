package main

import (
	"fmt"
	"handling-basic-request/handling"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handling.Hello)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
