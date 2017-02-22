package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/mail"

	"github.com/gorilla/mux"
)

var tpl *template.Template

type QuoteForm struct {
	Name, Email, Phone, CompanyName, Message string
}
type ContactForm struct {
	Name, Email, Phone, Message string
}

// // // // // // // //
//
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

		// execute email template
		t := template.New("send-q.gohtml")

		var err error

		t, err = t.ParseFiles("templates/send-q.gohtml")
		if err != nil {
			log.Println(err)
		}

		// convert email to string
		var tpl bytes.Buffer
		if err := t.Execute(&tpl, q); err != nil {
			log.Println(err)
		}

		// template string
		result := tpl.String()

		// prep email
		msg := &mail.Message{
			Sender:   "americanbrokerageapp@gmail.com",
			To:       []string{"Juliet imarchenko@gmail.com"},
			ReplyTo:  q.Email,
			Subject:  q.CompanyName + " has requested a quote from your website.",
			HTMLBody: result,
			// Attachments: []Attachment{
			// 	Attachment{
			// 		Name: "main.go",
			// 	},
			// },
		}

		// send email
		c := appengine.NewContext(r)
		if err := mail.Send(c, msg); err != nil {
			log.Fatalln(err)
		}
	}

}

// // // // // // // //
//
// company route handler
func company(w http.ResponseWriter, r *http.Request) {

	// execute tamplate
	err := tpl.ExecuteTemplate(w, "company.gohtml", nil)
	if err != nil {
		fmt.Println(err)
	}

}

// // // // // // // //
//
// company route handler
func ltl(w http.ResponseWriter, r *http.Request) {

	// execute tamplate
	err := tpl.ExecuteTemplate(w, "ltl.gohtml", nil)
	if err != nil {
		fmt.Println(err)
	}

}

// // // // // // // //
//
// terms route handler
func terms(w http.ResponseWriter, r *http.Request) {

	// execute tamplate
	err := tpl.ExecuteTemplate(w, "terms.gohtml", nil)
	if err != nil {
		fmt.Println(err)
	}

}

// // // // // // // //
//
// privacy route handler
func privacy(w http.ResponseWriter, r *http.Request) {

	// execute tamplate
	err := tpl.ExecuteTemplate(w, "privacy.gohtml", nil)
	if err != nil {
		fmt.Println(err)
	}

}

// // // // // // // //
//
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

}

//goapp deploy -application american-brokerage-158918 -version 1 .
