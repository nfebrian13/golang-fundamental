package main

import (
	"fmt"
	"net/http"
)

/* macam-macam handler
   1. NotFoundHandler
   2. RedirectHandler
   3. StripPrefix
   4. TimeoutHandler
   5. FileServer
*/

type myHandler struct {
	greeting string
}

func (mh myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("%v world", mh.greeting)))
}

func main() {

	/* cara pertama, handle request */
	// http.Handle("/", &myHandler{greeting: "Hello"})

	/* cara kedua, handle request */
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go!"))
	})

	http.ListenAndServe(":8000", nil)
}
