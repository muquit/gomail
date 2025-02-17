package main

import (
	"encoding/base64"
	"fmt"
	"github.com/muquit/gomail"
	"log"
	"os"
)

func createOauth2String(username, token string) string {
	authString := fmt.Sprintf("user=%s\001auth=Bearer %s\001\001", username, token)
	return base64.StdEncoding.EncodeToString([]byte(authString))
}

func main() {
	token := os.Getenv("SMTP_OAUTH_TOKEN")
	fromEmail := os.Getenv("SMTP_FROM_EMAIL")
	toEmail := os.Getenv("SMTP_TO_EMAIL")

	if token == "" {
		log.Fatal("SMTP_OAUTH_TOKEN environment variable is not set")
	}
	if fromEmail == "" {
		log.Fatal("SMTP_FROM_EMAIL environment variable is not set")
	}
	if toEmail == "" {
		log.Fatal("SMTP_TO_EMAIL environment variable is not set")
	}

	// Debug: Print exact same format as Python for comparison
	//log.Printf("Go auth string: %s", createOauth2String(fromEmail, "XXXX"))

	dialer := gomail.NewXOAuth2Dialer("smtp.gmail.com", 587, fromEmail, token)

	m := gomail.NewMessage()
	m.SetHeader("From", fromEmail)
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", "XOAUTH2 Test from gomail")
	m.SetBody("text/plain", "This is a test email using XOAUTH2 authentication")

	if err := dialer.DialAndSend(m); err != nil {
		log.Printf("Error details: %v", err)
		log.Fatalf("Error sending mail: %v", err)
	}
	log.Println("Mail sent successfully!")
}
