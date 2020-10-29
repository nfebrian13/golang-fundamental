package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Hello, Gopher")
	port := 3000
	isStarted := startWebServer2(port)
	fmt.Println(isStarted)

	fmt.Println()
	p, err := startWebServer1(port)
	fmt.Println(p, err)
}

func startWebServer(port int, numberOfRetries int) {
	fmt.Println("Starting server...")
	// do important things
	fmt.Println("Server started on port", port)
	fmt.Println("Number of retries", numberOfRetries)
}

func startWebServer1(port int) (int, error) {
	fmt.Println("Starting server...")
	// do important things
	fmt.Println("Server started on port", port)
	return port, errors.New("Something went wrong")
}

func startWebServer2(port int) bool {
	fmt.Println("Starting server...")
	// do important things
	fmt.Println("Server started on port", port)
	return true
}
