package email

import (
	"fmt"
	"github.com/danyaobertan/exchangemonitor/internal/config"
	"net/smtp"
)

// SendEmail sends an email using the specified SMTP configuration.
func SendEmail(smtpConfig config.SMTPConfig, recipient, subject, body string) error {
	// Setup SMTP client configuration.
	addr := fmt.Sprintf("%s:%s", smtpConfig.Host, smtpConfig.Port)
	auth := smtp.PlainAuth("", smtpConfig.Username, smtpConfig.Password, smtpConfig.Host)

	// Prepare the email message.
	message := []byte("To: " + recipient + "\r\n" +
		"From: " + smtpConfig.Username + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")

	// Send the email.
	err := smtp.SendMail(addr, auth, smtpConfig.Username, []string{recipient}, message)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
