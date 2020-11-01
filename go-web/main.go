package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type M map[string]interface{}

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

	http.HandleFunc("/", handlerHttpMethodRequest)

	// http.HandleFunc("/index", handlerRenderViewString)
	// http.HandleFunc("/", handlerRenderRedirectIndex)

	// http.HandleFunc("/", handlerRenderSpecific)
	// http.HandleFunc("/test", handlerRenderSpecificTest)

	// http.HandleFunc("/", handlerIndexSuperhero)
	// http.HandleFunc("/", handlerIndex)
	// http.HandleFunc("/about", handlerAbout)
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

/*
func handlerIndex(w http.ResponseWriter, r *http.Request) {
	// Fungsi Join ini digunakan untuk menggabungkan folder atau file
	// atau keduanya menjadi sebuah path

	var filepath = path.Join("views", "index.html")
	var tmpl, err = template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data = M{"name": "Batman"}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
} */

/* handlerIndex solusi render view menggunakan template.ParseGlob.
   kekurangan template.ParseGlob jika banyak views di dalam satu project dan semua view tersebut
   tidak terpakai atau terpakai sebagian maka parsing akan sia sia

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	var tmpl, err = template.ParseGlob("views/*")
	var data = M{"name": "Batman"}
	err = tmpl.ExecuteTemplate(w, "index", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
} */

/* handlerAbout solusi render view menggunakan template.ParseGlob.
   kekurangan template.ParseGlob jika banyak views di dalam satu project dan semua view tersebut
   tidak terpakai atau terpakai sebagian maka parsing akan sia sia

func handlerAbout(w http.ResponseWriter, r *http.Request) {
	var tmpl, err = template.ParseGlob("views/*")
	var data = M{"name": "Batman"}
	err = tmpl.ExecuteTemplate(w, "about", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
} */

/*
func handlerIndex(w http.ResponseWriter, r *http.Request) {
	var data = M{"name": "Batman"}
	var tmpl = template.Must(template.ParseFiles(
		"views/index.html",
		"views/_header.html",
		"views/_message.html",
	))
	var err = tmpl.ExecuteTemplate(w, "index", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handlerAbout(w http.ResponseWriter, r *http.Request) {
	var data = M{"name": "Batman"}
	var tmpl = template.Must(template.ParseFiles(
		"views/about.html",
		"views/_header.html",
		"views/_message.html",
	))
	var err = tmpl.ExecuteTemplate(w, "about", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
} */

/*
func handlerIndex(w http.ResponseWriter, r *http.Request) {
	var message = "Hello world from Nana!"
	w.Write([]byte(message))
} */

func handlerHello(w http.ResponseWriter, r *http.Request) {
	var message = "Hello world!"
	w.Write([]byte(message))
}

type Info struct {
	Affiliation string
	Address     string
}

type Person struct {
	Name    string
	Gender  string
	Hobbies []string
	Info    Info
}

func (t Info) GetAffiliationDetailInfo() string {
	return "have 31 divisions"
}

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	var person = Person{
		Name:    "Bruce Wayne",
		Gender:  "male",
		Hobbies: []string{"Reading Books", "Traveling", "Buying things"},
		Info:    Info{"Wayne Enterprises", "Gotham City"},
	}

	var tmpl = template.Must(template.ParseFiles("views/view.html"))
	if err := tmpl.Execute(w, person); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type Superhero struct {
	Name    string
	Alias   string
	Friends []string
}

func (s Superhero) SayHello(from string, message string) string {
	return fmt.Sprintf("%s said: \"%s\"", from, message)
}

func handlerIndexSuperhero(w http.ResponseWriter, r *http.Request) {
	var person = Superhero{
		Name:    "Bruce Wayne",
		Alias:   "Batman",
		Friends: []string{"Superman", "Flash", "Green Lantern"},
	}

	var tmpl = template.Must(template.ParseFiles("views/view_batman.html"))
	if err := tmpl.Execute(w, person); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handlerRenderSpecific(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.New("index").ParseFiles("views/view_render_specific.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handlerRenderSpecificTest(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.New("test").ParseFiles("views/view_render_specific.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

const (
	view string = `<html>
    <head>
        <title>Template</title>
    </head>
    <body>
        <h1>Hello</h1>
    </body>
	</html>`
)

func handlerRenderViewString(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.New("main-template").Parse(view))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func handlerRenderRedirectIndex(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/index", http.StatusTemporaryRedirect)

}

/* handler http method */
func handlerHttpMethodRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		w.Write([]byte("post"))
	case "GET":
		w.Write([]byte("get"))
	default:
		http.Error(w, "", http.StatusBadRequest)
	}
}
