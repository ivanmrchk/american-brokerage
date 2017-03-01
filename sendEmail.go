package main

import (
	"bytes"
	"html/template"
	"log"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/mail"
)

type QuoteForm struct {
	Name, Email, Phone, CompanyName, Message string
}
type ContactForm struct {
	Name, Email, Phone, Message, Subject string
}
type QuotePageForm struct {
	Name           string
	Title          string
	CompanyName    string
	CompanyAddress string
	City           string
	State          string
	Zipcode        string
	Email          string
	ConfirmEamil   string
	Phone          string
	Pickup         string
	Drop           string
	Message        string
}

func (q QuoteForm) SendQuoteForm(w http.ResponseWriter, r *http.Request) {
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
		To:       []string{"Jack imarchenko@gmail.com"},
		ReplyTo:  q.Email,
		Subject:  q.Name,
		HTMLBody: result,
	}

	// send email
	x := appengine.NewContext(r)
	if err := mail.Send(x, msg); err != nil {
		log.Fatalln(err)
	}
}
func (qp QuotePageForm) SendQuotePage(w http.ResponseWriter, r *http.Request) {
	// execute email template
	t := template.New("send-qp.gohtml")

	var err error

	t, err = t.ParseFiles("templates/send-qp.gohtml")
	if err != nil {
		log.Println(err)
	}

	// convert email to string
	var tpl bytes.Buffer
	if err := t.Execute(&tpl, qp); err != nil {
		log.Println(err)
	}

	// template string
	result := tpl.String()

	// prep email
	msg := &mail.Message{
		Sender:   "americanbrokerageapp@gmail.com",
		To:       []string{"Jack imarchenko@gmail.com"},
		ReplyTo:  qp.Email,
		Subject:  qp.Name + " has requested a quote.",
		HTMLBody: result,
	}

	// send email
	x := appengine.NewContext(r)
	if err := mail.Send(x, msg); err != nil {
		log.Fatalln(err)
	}
}

func (c ContactForm) SendContactForm(w http.ResponseWriter, r *http.Request) {
	// execute email template
	t := template.New("send-c.gohtml")

	var err error

	t, err = t.ParseFiles("templates/send-c.gohtml")
	if err != nil {
		log.Println(err)
	}

	// convert email to string
	var tpl bytes.Buffer
	if err := t.Execute(&tpl, c); err != nil {
		log.Println(err)
	}

	// template string
	result := tpl.String()

	// prep email
	msg := &mail.Message{
		Sender:   "americanbrokerageapp@gmail.com",
		To:       []string{"jack imarchenko@gmail.com"},
		ReplyTo:  c.Email,
		Subject:  c.Name + " is interested in: " + c.Subject,
		HTMLBody: result,
	}

	// send email
	x := appengine.NewContext(r)
	if err := mail.Send(x, msg); err != nil {
		log.Fatalln(err)
	}
}
