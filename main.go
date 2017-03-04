package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var tpl *template.Template

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

		// quote form data
		q := QuoteForm{
			Name:        r.FormValue("qname"),
			Email:       r.FormValue("qemail"),
			Phone:       r.FormValue("qphone"),
			CompanyName: r.FormValue("qcompanyName"),
			Message:     r.FormValue("qmessage"),
		}

		q.SendQuoteForm(w, r)
	}

}

// contactus route handler
func contactUs(w http.ResponseWriter, r *http.Request) {

	// parse form value
	r.ParseForm()

	// execute tamplate
	err := tpl.ExecuteTemplate(w, "contact-us.gohtml", nil)

	if err != nil {
		fmt.Println(err)
	}

	if r.Method == http.MethodPost {

		// quote form data
		c := ContactForm{
			Name:    r.FormValue("name"),
			Email:   r.FormValue("email"),
			Phone:   r.FormValue("phone"),
			Subject: r.FormValue("subject"),
			Message: r.FormValue("message"),
		}

		c.SendContactForm(w, r)
	}

}

//  drive  route handler
func drive(w http.ResponseWriter, r *http.Request) {

	// execute tamplate
	err := tpl.ExecuteTemplate(w, "drive.gohtml", nil)
	if err != nil {
		fmt.Println(err)
	}

}

// quote route handler
func quote(w http.ResponseWriter, r *http.Request) {

	// parse form value
	r.ParseForm()
	// execute tamplate
	err := tpl.ExecuteTemplate(w, "quote.gohtml", nil)
	if err != nil {
		fmt.Println(err)
	}
	// send Form
	if r.Method == http.MethodPost {

		// quote form data
		qp := QuotePageForm{
			Name:           r.FormValue("name"),
			Title:          r.FormValue("title"),
			CompanyName:    r.FormValue("companyName"),
			CompanyAddress: r.FormValue("companyAddress"),
			City:           r.FormValue("city"),
			State:          r.FormValue("state"),
			Zipcode:        r.FormValue("zip"),
			Email:          r.FormValue("email"),
			ConfirmEamil:   r.FormValue("confirmEmail"),
			Phone:          r.FormValue("phone"),
			Pickup:         r.FormValue("pickup"),
			Drop:           r.FormValue("drop"),
			Message:        r.FormValue("message"),
		}

		qp.SendQuotePage(w, r)
	}
}

// carriers route handler
func carriers(w http.ResponseWriter, r *http.Request) {

	// parse form value
	r.ParseForm()
	// execute tamplate
	err := tpl.ExecuteTemplate(w, "carriers.gohtml", nil)
	if err != nil {
		fmt.Println(err)
	}
	// send Form
	if r.Method == http.MethodPost {

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

// less then truck load route handler
func ltl(w http.ResponseWriter, r *http.Request) {

	// execute tamplate
	err := tpl.ExecuteTemplate(w, "ltl.gohtml", nil)
	if err != nil {
		fmt.Println(err)
	}

}

// less then truck load route handler
func fullLoad(w http.ResponseWriter, r *http.Request) {

	// execute tamplate
	err := tpl.ExecuteTemplate(w, "full-load.gohtml", nil)
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

func init() {

	// parse template dir
	tpl = template.Must(template.ParseGlob("templates/*"))

	// router vars
	r := mux.NewRouter()
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))

	// routes
	r.HandleFunc("/", index)
	r.HandleFunc("/company", company)
	r.HandleFunc("/service", ltl)
	r.HandleFunc("/ltl", ltl)
	r.HandleFunc("/full-load", fullLoad)
	r.HandleFunc("/privacy", privacy)
	r.HandleFunc("/terms", terms)
	r.HandleFunc("/contact-us", contactUs)
	r.HandleFunc("/drive", drive)
	r.HandleFunc("/quote", quote)
	r.HandleFunc("/drive/carriers", carriers)

	// Not found
	r.NotFoundHandler = http.HandlerFunc(customNotFound)

	//http.ListenAndServe(":8080", r)

	http.Handle("/", r)

}

//goapp deploy -application american-brokerage-158918 -version 1 .
