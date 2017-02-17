package main

import (
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/julienschmidt/httprouter"
)

var tpl *template.Template

func main() {

	// parse template dir
	tpl = template.Must(template.ParseGlob("templates/*"))

	// set new router
	//mux := httprouter.New()

	// define routes
	http.HandleFunc("/", index)

	//mux.ServeFiles("/public", http.Dir(""))
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./public"))))

	http.ListenAndServe(":8080", nil)

}

// root route handler
func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	// execute tamplate
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		fmt.Println(err)
	}

}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		err := tpl.ExecuteTemplate(w, "404.gohtml", nil)
		if err != nil {
			fmt.Println(err)
		}
	}

}

//goapp deploy -application american-brokerage-158918 -version 1 .
