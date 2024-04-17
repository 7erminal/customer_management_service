package functions

import (
	"net/smtp"

	"github.com/beego/beego/v2/core/logs"
)

func SendEmail(username string) {
	auth := smtp.PlainAuth("", "bede.abbe@solisfinance.com", "dqcmwcuyvfjwcuyd", "smtp.gmail.com")

	// Here we do it all: connect to our server, set up a message and send it

	to := []string{"bede.abbe91@gmail.com"}

	msg := []byte("To: bede.abbe91@gmail.com\r\n" +

		"Subject: Why aren’t you using Mailtrap yet?\r\n" +

		"\r\n" +

		"Here’s the space for our great sales pitch\r\n")

	err := smtp.SendMail("smtp.gmail.com:587", auth, "bede.abbe@solisfinance.com", to, msg)

	if err != nil {

		logs.Debug(err)

	}
}
