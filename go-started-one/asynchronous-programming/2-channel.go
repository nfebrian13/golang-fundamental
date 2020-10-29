package main

import (
	"fmt"
	"runtime"
)

func main() {

	runtime.GOMAXPROCS(8)

	ch := make(chan string)
	doneCh := make(chan bool)

	go abcEng(ch)
	go printer(ch, doneCh)

	fmt.Println("this execute firts")

	<-doneCh

}

func abcEng(ch chan string) {
	for i := byte('a'); i <= byte('z'); i++ {
		ch <- string(i)
	}
	close(ch)
}

func printer(ch chan string, doneCh chan bool) {
	for i := range ch {
		println(i)
	}
	doneCh <- true
}
