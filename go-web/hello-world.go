package hello-world

import (
	"fmt"
	"net/http"
)

func helloworld() {
	/* HandleFunc, Fungsi ini digunakan untuk routing,
	menentukan aksi dari sebuah url tertentu ketika diakses
	(di sini url tersebut kita sebut sebagai rute/route). */
	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/index", handlerIndex)
	http.HandleFunc("/hello", handlerHello)

	var address = "localhost:9000"
	fmt.Printf("server started at %s\n", address)
	/* digunakan membuat sekaligus start server baru,
	   dengan parameter pertama adalah alamat web server yang diiginkan
	   (bisa diisi host, host & port, atau port saja). */
	err := http.ListenAndServe(address, nil) // Fungsi ini digunakan untuk membuat web server baru
	if err != nil {
		fmt.Println(err.Error())
	}
}

/* Route handler atau handler atau parameter kedua fungsi http.HandleFunc(),
   adalah sebuah fungsi dengan ber-skema func (ResponseWriter, *Request). */

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	var message = "Hello world from Nana!"
	w.Write([]byte(message))
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	var message = "Hello world!"
	w.Write([]byte(message))
}
