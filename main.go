package main

import (
	"fmt"
	"net/smtp"
	"os"
)

func main() {
	var args []string = os.Args[1:]

	from := args[0]
	password := args[1]

	to := []string{
		args[2],
	}

	//Config
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message
	message := []byte(args[3])

	// Auth sender
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Send Email
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Email has been sent!")

}
