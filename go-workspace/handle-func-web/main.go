package main

import (
	"fmt"
	"net/http"
)

type myHandler struct {
	greeting string
}

func (mh myHandler) ServerHttp(w http.ResponseWriter, r *http.Request){
	w.Write([]byte(fmt.Sprintf("%v world",mh.greeting)))
}

func main(){
	// http.Handle("/", &myHandler{greeting:"hello"})
    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Hello World!"))
	})
	http.ListenAndServe(":8000",nil)
}
 