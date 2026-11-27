package functions

import (
	"context"
	"customer_management_service/internal/trace"
	"net/smtp"
)

func SendEmail(ctx context.Context, email string, subject_ string, message string) {
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

		trace.Logger(ctx, trace.Service).Error("email delivery failed")

	}
}
