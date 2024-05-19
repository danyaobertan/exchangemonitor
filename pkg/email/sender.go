package email

import (
	"bytes"
	"fmt"
	"github.com/danyaobertan/exchangemonitor/internal/config"
	"github.com/danyaobertan/exchangemonitor/models"
	"html/template"
	"net/smtp"
)

// SendEmail sends an email using the specified SMTP configuration.
func SendEmail(smtpConfig config.SMTPConfig, recipient string, object models.EmailDataObject) error {
	tmpl, err := template.ParseFiles("templates/email_template.html")
	if err != nil {
		return fmt.Errorf("error parsing email template: %w", err)
	}

	// Execute the template with the data
	var body bytes.Buffer
	if err := tmpl.Execute(&body, object); err != nil {
		return fmt.Errorf("error executing email template: %w", err)
	}

	// Setup SMTP client configuration.
	addr := fmt.Sprintf("%s:%s", smtpConfig.Host, smtpConfig.Port)
	auth := smtp.PlainAuth("", smtpConfig.Username, smtpConfig.Password, smtpConfig.Host)

	// Prepare the email message with HTML content.
	message := []byte("To: " + object.Name + "\r\n" +
		"From: " + smtpConfig.Username + "\r\n" +
		"Subject: " + object.Subject + "\r\n" +
		"MIME-Version: 1.0;\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\";\r\n" +
		"\r\n" +
		body.String() + "\r\n")

	// Send the email.
	err = smtp.SendMail(addr, auth, smtpConfig.Username, []string{recipient}, message)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
