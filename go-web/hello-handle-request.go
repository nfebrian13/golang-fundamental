package main

import (
	"fmt"
	"net/http"
)

func main() {
	/* HandleFunc, Fungsi ini digunakan untuk routing,
	menentukan aksi dari sebuah url tertentu ketika diakses
	(di sini url tersebut kita sebut sebagai rute/route). */
	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/index", handlerIndex)
	http.HandleFunc("/hello", handlerHello)

	var address = ":9001"
	fmt.Printf("server started at %s\n", address)

	server := new(http.Server)
	server.Addr = address
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	}
}
func handlerIndex(w http.ResponseWriter, r *http.Request) {
	var message = "Hello world from Nana!"
	w.Write([]byte(message))
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	var message = "Hello world!"
	w.Write([]byte(message))
}
