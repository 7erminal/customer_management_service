package functions

import (
	"net/smtp"

	"github.com/beego/beego/v2/core/logs"
)

func SendEmail(email string, link string) {
	// Create app password in gmail to use here
	auth := smtp.PlainAuth("", "bede.abbe@gmail.com", "psxglveajilrvisa", "smtp.gmail.com")

	// Here we do it all: connect to our server, set up a message and send it

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject := "Subject: AMC Sales & Rentals user invitation\n"
	body := "<html><body><table></table><tr><th>AMC Sales & Rentals</th></tr><tr></tr><tr><td>User invitation</td></tr><tr><td>" + link + "</td></tr><tr></tr><tr><a href='" + link + "'>Accept Invite</a></tr></body></html>"

	to := []string{email}

	msg := []byte(subject + mime + body)

	err := smtp.SendMail("smtp.gmail.com:587", auth, "bede.abbe@gmail.com", to, msg)

	if err != nil {

		logs.Debug(err)

	}
}
