package main

import (
	"fmt"
	"net/smtp"
)

// rename mail2  to mail for actual mail.... def mail func for testing ...
func mail2(a string, b string) {

	// Sender data.
	from := "me@me.me"
	password := ""

	// Receiver email address.
	to := []string{
		a,
	}

	// smtp server configuration.
	smtpHost := "smtp.cse.iitk.ac.in"
	smtpPort := "587"
	for len(b) != 4 {
		b = "0" + b
	}
	// Message.
	message := []byte("OTP: " + b)

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	//   fmt.Println("Email Sent Successfully!")
}

func mail(a string, b string) {

	fmt.Println(a, b)
}
