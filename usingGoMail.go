package main

import (
	"bytes"
	"log"
	"os"
	"text/template"

	"gopkg.in/gomail.v2"
)

func SendGoMail(templatePath string){
	//Get HTML
	var body bytes.Buffer
	t, err := template.ParseFiles(templatePath)
	if err != nil{
		log.Fatal(err)
	}
	t.Execute(&body, struct{Name string}{Name: "Agodzo"})


	//send with gomail
	m := gomail.NewMessage()
	m.SetHeader("From", "fahdmoh.1@gmail.com")
	m.SetHeader("To", "fahdmoh.1@gmail.com")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", body.String())
	m.Attach("./sourcecode.png")

	d := gomail.NewDialer("smtp.gmail.com", 587, "fahdmoh.1@gmail.com", os.Getenv("PASSWORD"))

	// Send the email to myself.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}


}