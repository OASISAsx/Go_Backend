package infrastructure

import (
	"os"

	"github.com/joho/godotenv"
)

func GmailPassword() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("GMAIL_PASSWORD")
}
func GmailUser() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("GMAIL_USER")
}
func SmtpAddr() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("SMTP_ADDR")
}
func SmtpFrom() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("SMTP_FROM")
}
func SmtpHost() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("SMTP_HOST")
}
