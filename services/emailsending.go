package services

import "net/smtp"

func SendEmail(to string, message string) error {
	// Sender data.
	from := "wijnbergenmark@gmail.com"
	password := "<Email Password>"

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	body := []byte(message)

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	return smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, body)
}
