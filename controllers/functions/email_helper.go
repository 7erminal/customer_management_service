package functions

import (
	"context"
	"net/smtp"

	"customer_management_service/logging"
)

// SendEmail sends the invite e-mail in the background. ctx carries the
// tracing identifiers of the request that spawned it (use
// context.WithoutCancel so the mail still goes out after the request
// finishes); the IDs are only used for logging, never for auth.
func SendEmail(ctx context.Context, email string, subject_ string, message string) {
	log := logging.For(ctx)
	// Create app password in gmail to use here
	// auth := smtp.PlainAuth("", "bede.abbe@gmail.com", "psxglveajilrvisa", "smtp.gmail.com")
	auth := smtp.PlainAuth("", "info@amcrentalsgh.com", "eeastbiqnvvjzyjb", "smtp.gmail.com")

	// Here we do it all: connect to our server, set up a message and send it

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject := "Subject:" + subject_
	body := message

	to := []string{email}

	msg := []byte(subject + mime + body)

	err := smtp.SendMail("smtp.gmail.com:587", auth, "info@amcrentalsgh.com", to, msg)

	if err != nil {
		log.Debug("sending e-mail failed", "error", err.Error())
	}
}
