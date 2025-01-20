package email

import "gopkg.in/gomail.v2"

type mailer struct {
	dialer *gomail.Dialer
	branch string
}

type Mailer interface {
	Send(emailToArr []string, title string, body string)
}
