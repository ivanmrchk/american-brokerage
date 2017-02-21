package main

import (
	"bytes"
	"html/template"
	"log"

	gomail "gopkg.in/gomail.v2"
)

type QuoteForm struct {
	Name        string
	Email       string
	Phone       string
	CompanyName string
	Message     string
}

// send quote to email
func (q QuoteForm) SendQuote() {

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

	m.SetHeaders(map[string][]string{
		"From":     {m.FormatAddress(q.Email, q.Name)},
		"To":       {m.FormatAddress("imarchenko@gmail.com", "Ivan")},
		"Cc":       {m.FormatAddress("mericanbrokerageapp@gmail.com", "Admin")},
		"Reply-to": {m.FormatAddress("marchenkoi@outlook.com", "Jack")},
		"Subject":  {q.Name + " Has requested a quote fon your website."},
	})
	// m.SetHeader("From", q.Email)
	// m.SetHeader("To", "imarchenko@gmail.com")
	// m.SetAddressHeader("Cc", "americanbrokerageapp@gmail.com", "Admin")
	// m.SetHeader("Subject", q.Name+" Has requested a quote fon your website.")
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
