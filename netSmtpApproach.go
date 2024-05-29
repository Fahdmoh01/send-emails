package main

import (
	"bytes"
	"log"
	"net/smtp"
	"os"
	"text/template"
)

func SendMailSimple(subject string, body string, to []string){
	auth :=smtp.PlainAuth(
		"",
		"fahdmoh.1@gmail.com",
		os.Getenv("PASSWORD"),
		"smtp.gmail.com",
	)

	msg := "Subject: " + subject + "\n" + body

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"fahdmoh.1@gmail.com",
		to,
		[]byte(msg),
	)

	if err != nil{
		log.Fatal(err)
	}
}

func SendMailSimpleHTML(subject string, templatePath string, to []string){
	//Get HTML
	var body bytes.Buffer
	t, err := template.ParseFiles(templatePath)
	if err != nil{
		log.Fatal(err)
	}

	t.Execute(&body, struct{Name string}{Name: "Agodzo"})

	auth :=smtp.PlainAuth(
		"",
		"fahdmoh.1@gmail.com",
		os.Getenv("PASSWORD"),
		"smtp.gmail.com",
	)

	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-\";"
	msg := "Subject: " + subject + "\n" + headers + "\n\n" +body.String()

	err = smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"fahdmoh.1@gmail.com",
		to,
		[]byte(msg),
	)

	if err != nil{
		log.Fatal(err)
	}
}

