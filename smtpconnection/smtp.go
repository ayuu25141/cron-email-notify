package smtpconnection

import (
	"fmt"
	"net/smtp"
	"os"
)

type SMTPConfig struct {
	Server string
	Port   string
	Login  string
	APIKey string
	From   string
	Auth   smtp.Auth
}

// NewBrevoSMTPConfig initializes Brevo SMTP config from env
func NewBrevoSMTPConfig() (*SMTPConfig, error) {
	server := os.Getenv("Smtpserver")
	port := os.Getenv("Smtpport")
	login := os.Getenv("Smtplogin")
	apiKey := os.Getenv("Smtpapi")
	from := os.Getenv("Smtpfrom")

	if server == "" || port == "" || login == "" || apiKey == "" || from == "" {
		return nil, fmt.Errorf("missing Brevo SMTP environment variables")
	}

	auth := smtp.PlainAuth(
		"",
		login,
		apiKey,
		server,
	)

	return &SMTPConfig{
		Server: server,
		Port:   port,
		Login:  login,
		APIKey: apiKey,
		From:   from,
		Auth:   auth,
	}, nil
}
