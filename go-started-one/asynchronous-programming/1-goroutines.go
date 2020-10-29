package goroutine

import (
	"fmt"
	"runtime"
	"time"
)

func goroutine() {

	/*	runtime.GOMAXPROCS(8)

		go abcEng()
		fmt.Println("this execute firts")
		time.Sleep(100 * time.Millisecond) */

	runtime.GOMAXPROCS(1)
	go printHello(10, "Hello World")
	printHello(10, "Hello Nana")

	fmt.Println("this execute firts")
	time.Sleep(100 * time.Millisecond)
}

func abcEng() {
	for i := byte('a'); i <= byte('z'); i++ {
		fmt.Println(string(i))
	}
}

func printHello(till int, word string) {
	for i := 0; i < till; i++ {
		fmt.Println(i, " ", word)
	}
}
