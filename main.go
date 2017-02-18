package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var tpl *template.Template

func main() {

	// parse template dir
	tpl = template.Must(template.ParseGlob("templates/*"))

	// set new router
	mux := httprouter.New()

	// define routes
	mux.GET("/", index)
	mux.POST("/", postIndex)
	mux.GET("/terms", terms)
	mux.GET("/privacy", privacy)
	mux.NotFound = http.HandlerFunc(customNotFound)

	// serving static files
	mux.ServeFiles("/assets/*filepath", http.Dir("assets"))
	http.ListenAndServe(":8080", mux)

}

// root route handler
func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	// execute tamplate
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)

	if err != nil {
		fmt.Println(err)
	}

}
func postIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("a")
	// execute tamplate
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)

	if err != nil {
		fmt.Println(err)
	}

}

// terms route handler
func terms(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	// execute tamplate
	err := tpl.ExecuteTemplate(w, "terms.gohtml", nil)
	if err != nil {
		fmt.Println(err)
	}

}

// privacy route handler
func privacy(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	// execute tamplate
	err := tpl.ExecuteTemplate(w, "privacy.gohtml", nil)
	if err != nil {
		fmt.Println(err)
	}

}

// custom 404 route hanler
func customNotFound(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path, "not found route")

	err := tpl.ExecuteTemplate(w, "404.gohtml", nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(r.Method)

}

//goapp deploy -application american-brokerage-158918 -version 1 .
