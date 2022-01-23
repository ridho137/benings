package util

import (
	"benings/model"
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"

	gomail "gopkg.in/mail.v2"
)

func EmailSender(w http.ResponseWriter, destination string, tittle string, info model.Info) {
	log.Println(info)
	t, _ := template.ParseFiles("util/template.html")
	var tpl bytes.Buffer
	fmt.Println(t.Execute(&tpl, info))

	result := tpl.String()

	email := gomail.NewMessage()
	from := "otp.benings@gmail.com"
	password := "B*nings_"
	to := destination
	subject := tittle
	body := result
	host := "smtp.gmail.com"
	port := 587

	email.SetHeader("From", from)
	email.SetHeader("To", to)
	email.SetHeader("Subject", subject)
	email.SetBody("text/html", body)

	send := gomail.NewDialer(host, port, from, password)
	if err := send.DialAndSend(email); err != nil {
		fmt.Println(err)
		panic(err)
	}
}
