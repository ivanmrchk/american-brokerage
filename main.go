package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"

	gomail "gopkg.in/gomail.v2"

	"github.com/julienschmidt/httprouter"
)

var tpl *template.Template

type quoteForm struct {
	Name    string
	Email   string
	Phone   string
	Subject string
	Message string
}

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

	r.ParseForm()

	// execute tamplate
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)

	if err != nil {
		fmt.Println(err)
	}

	// quote form data
	q := quoteForm{Name: r.FormValue("qname"),
		Email:   r.FormValue("qemail"),
		Phone:   r.FormValue("qphone"),
		Subject: r.FormValue("qsubject"),
		Message: r.FormValue("qmessage"),
	}

	// grab Name from q
	fromN := q.Name
	fromE := q.Email

	// send email
	q.sendMail(fromN, fromE)

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

// custom 404 route handler
func customNotFound(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path, "not found route")

	err := tpl.ExecuteTemplate(w, "404.gohtml", nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(r.Method)

}

// send quote to email
func (q quoteForm) sendMail(fromName string, fromEmail string) {

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
	result := tpl.String()

	// prepare email
	m := gomail.NewMessage()
	m.SetHeader("From", fromEmail)
	m.SetHeader("To", "americanbrokerageapp@gmail.com")
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", fromName+" Has requested a quote fon your website.")
	m.SetBody("text/html", result)
	//m.Attach("template.html")

	// prepare dialer
	d := gomail.NewDialer("smtp.gmail.com", 587, "americanbrokerageapp@gmail.com", "M)FrFC=pwR8(<?#B")

	// send email
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	log.Println("mail sent", q)
}

//goapp deploy -application american-brokerage-158918 -version 1 .
