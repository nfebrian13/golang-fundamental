package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Superhero struct {
	Name    string
	Alias   string
	Friends []string
}

func (s Superhero) SayHello(from string, message string) string {
	return fmt.Sprintf("%s said: \"%s\"", from, message)
}

func HandlerIndexSuperhero(w http.ResponseWriter, r *http.Request) {
	var person = Superhero{
		Name:    "Bruce Wayne",
		Alias:   "Batman",
		Friends: []string{"Superman", "Flash", "Green Lantern"},
	}

	var tmpl = template.Must(template.ParseFiles("view_batman.html"))
	if err := tmpl.Execute(w, person); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
