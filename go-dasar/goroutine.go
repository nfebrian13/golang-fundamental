package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(2)
	go sayHello()

	time.Sleep(2 * time.Second)
	fmt.Printf("Current Unix Time: %v\n", time.Now().Unix())

}

func sayHello() {
	for i := 0; i < 10; i++ {
		fmt.Println(i, " Hello World")
	}
}
