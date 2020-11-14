package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/",handlerHello)
	http.HandleFunc("/hello",handlerHello)
	http.HandleFunc("/golang",handlerIndex)

	address := "localhost:9000"
    fmt.Printf("server started at %s\n", address)
    err := http.ListenAndServe(address, nil)
    if err != nil {
        fmt.Println(err.Error())
    }
}

func handlerIndex(w http.ResponseWriter, r *http.Request){
	message := "Hello golang web app!"
	w.Write([]byte(message))
}

func handlerHello(w http.ResponseWriter, r *http.Request){
	message := "Hello World!"
	w.Write([]byte(message))
}