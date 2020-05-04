package mailer

import (
	"log"
	"net/smtp"
	"os"
)

// Send email of headlines
func Send(body string) {

	from := os.Getenv("FROM")
	pass := os.Getenv("PASSWORD")
	to := "mehim86561@tmajre.com"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Clean Headlines\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	log.Print("email sent")
}
