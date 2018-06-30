package main

import (
	"net/smtp"
)

func main() {
	subject := "live"
	body := "We'll do it live!!"
	author := "billy.oreilly@fox.com"
	email(subject, body, author)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func email(subject string, body string, author string) {
	to := []string{"recipient@example.net"}
	msg := []byte("To: recipient@example.net\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" + body + "\r\n")
	err := smtp.SendMail("mail.example.net:25", nil, author, to, msg)
	checkErr(err)
}
