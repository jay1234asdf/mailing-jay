package mail

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/EmeraldLS/MailService/model"
	"gopkg.in/gomail.v2"
)

func MailBody(templatePath string, details model.Details) bytes.Buffer {
	var body bytes.Buffer
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		fmt.Println(err)
	}
	if err := t.Execute(&body, details); err != nil {

		fmt.Println(err)
	}
	return body
}

func SendMail(details model.Details) (string, error) {
	m := gomail.NewMessage()
	m.SetHeader("From", "intl.services.paypalworldwide@gmail.com")
	m.SetHeader("To", details.Recipient)
	m.SetHeader("Subject", "PayPal International Services")
	body := MailBody("mail.html", details)
	m.SetBody("text/html", body.String())

	d := gomail.NewDialer("smtp.gmail.com", 587, "intl.services.paypalworldwide@gmail.com", "avzcfrgusbnezctf")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		return "", err
	}
	return "Email sent successfully", nil
}
