package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
)

func main() {
	startWebApps()
}

func startWebApps() {

	/* contoh langsung penulisan dari handlerIndex
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	        // not yet implemented
	    })
	*/

	/* HandleFunc, Fungsi ini digunakan untuk routing,
	menentukan aksi dari sebuah url tertentu ketika diakses
	(di sini url tersebut kita sebut sebagai rute/route). */
	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/index", handlerIndex)
	http.HandleFunc("/hello", handlerHello)

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))

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
	/* Fungsi Join ini digunakan untuk menggabungkan folder atau file
	   atau keduanya menjadi sebuah path */
	var filepath = path.Join("views", "index.html")
	var tmpl, err = template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data = map[string]interface{}{
		"title": "Learning Golang Web",
		"name":  "Batman",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

/*
func handlerIndex(w http.ResponseWriter, r *http.Request) {
	var message = "Hello world from Nana!"
	w.Write([]byte(message))
} */

func handlerHello(w http.ResponseWriter, r *http.Request) {
	var message = "Hello world!"
	w.Write([]byte(message))
}
