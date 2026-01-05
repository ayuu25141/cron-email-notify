package maildata

import (
	"context"
	
	"log"
	"net/smtp"
	"time"

	"cronproject/dbconnection"
	"cronproject/smtpconnection"
)

// SendMailToAllUsers fetches emails from users table and sends mail
func SendMailToAllUsers(subject, body string) error {

	// 1️⃣ Get SMTP config
	smtpCfg, err := smtpconnection.NewBrevoSMTPConfig()
	if err != nil {
		return err
	}

	// 2️⃣ Query emails using existing pgx pool
	rows, err := dbconnection.Pool.Query(
		context.Background(),
		`SELECT email FROM users WHERE email IS NOT NULL`,
	)
	if err != nil {
		return err
	}
	defer rows.Close()

	// 3️⃣ Send emails (rate-limited)
	for rows.Next() {
		var email string
		if err := rows.Scan(&email); err != nil {
			continue
		}

		if err := sendSingleMail(smtpCfg, email, subject, body); err != nil {
			log.Println("email failed:", email, err)
			continue
		}

		// 🚨 important: avoid SMTP ban
		time.Sleep(1 * time.Second)
	}

	return nil
}

// sendSingleMail sends email to one recipient
func sendSingleMail(cfg *smtpconnection.SMTPConfig, to, subject, body string) error {

	msg := []byte(
		"From: " + cfg.From + "\r\n" +
			"To: " + to + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"MIME-Version: 1.0\r\n" +
			"Content-Type: text/html; charset=UTF-8\r\n\r\n" +
			body,
	)

	return smtp.SendMail(
		cfg.Server+":"+cfg.Port,
		cfg.Auth,
		cfg.From,
		[]string{to},
		msg,
	)
}
