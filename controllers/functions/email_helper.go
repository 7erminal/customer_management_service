package functions

import (
	"net/smtp"

	"github.com/beego/beego/v2/core/logs"
)

func SendEmail(email string, link string) {
	// Create app password in gmail to use here
	auth := smtp.PlainAuth("", "bede.abbe@gmail.com", "psxglveajilrvisa", "smtp.gmail.com")

	// Here we do it all: connect to our server, set up a message and send it

	to := []string{email}

	msg := []byte("To: " + email + "\r\n" +

		"Subject: Please click on the link below.\r\n" +

		"\r\n" +

		link + ".\r\nThis link will expire in 4 hours.\r\n")

	err := smtp.SendMail("smtp.gmail.com:587", auth, "bede.abbe@gmail.com", to, msg)

	if err != nil {

		logs.Debug(err)

	}
}
