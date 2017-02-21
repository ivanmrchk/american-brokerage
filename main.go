package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var tpl *template.Template

func main() {

	// parse template dir
	tpl = template.Must(template.ParseGlob("templates/*"))

	// router vars
	r := mux.NewRouter()
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))

	// routes
	// GET
	r.HandleFunc("/", index)
	r.HandleFunc("/company", company)
	r.HandleFunc("/service", ltl)
	r.HandleFunc("/ltl", ltl)
	r.HandleFunc("/privacy", privacy)
	r.HandleFunc("/terms", terms)

	// Not found
	r.NotFoundHandler = http.HandlerFunc(customNotFound)

	http.Handle("/", r)

	http.ListenAndServe(":8080", nil)
	fmt.Println("listening on port 8080")

}

// root route handler
func index(w http.ResponseWriter, r *http.Request) {

	// parse form value
	r.ParseForm()

	// execute tamplate
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)

	if err != nil {
		fmt.Println(err)
	}

	if r.Method == http.MethodPost {

		if err != nil {
			fmt.Println(err)
		}

		// quote form data
		q := QuoteForm{Name: r.FormValue("qname"),
			Email:       r.FormValue("qemail"),
			Phone:       r.FormValue("qphone"),
			CompanyName: r.FormValue("qcompanyName"),
			Message:     r.FormValue("qmessage"),
		}

		// send email
		q.SendQuote()
	}

}

// company route handler
func company(w http.ResponseWriter, r *http.Request) {

	// execute tamplate
	err := tpl.ExecuteTemplate(w, "company.gohtml", nil)
	if err != nil {
		fmt.Println(err)
	}

}

// company route handler
func ltl(w http.ResponseWriter, r *http.Request) {

	// execute tamplate
	err := tpl.ExecuteTemplate(w, "ltl.gohtml", nil)
	if err != nil {
		fmt.Println(err)
	}

}

// terms route handler
func terms(w http.ResponseWriter, r *http.Request) {

	// execute tamplate
	err := tpl.ExecuteTemplate(w, "terms.gohtml", nil)
	if err != nil {
		fmt.Println(err)
	}

}

// privacy route handler
func privacy(w http.ResponseWriter, r *http.Request) {

	// execute tamplate
	err := tpl.ExecuteTemplate(w, "privacy.gohtml", nil)
	if err != nil {
		fmt.Println(err)
	}

}

// custom 404 route handler
func customNotFound(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path, "not found route")

	err := tpl.ExecuteTemplate(w, "404.gohtml", nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(r.Method)

}

//goapp deploy -application american-brokerage-158918 -version 1 .
