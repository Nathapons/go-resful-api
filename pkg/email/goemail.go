package email

import (
	"crypto/tls"
	"strings"

	"github.com/go-resful-api/internal/configs"
	"gopkg.in/gomail.v2"
)

var defaultToEmail, defaultFromEmail string

func NewMailProtocol(cfg *configs.MailProtocolConfig, branch string) Mailer {
	dialer := gomail.NewDialer(cfg.MailerHost, cfg.MailerPort, cfg.MailerUser, cfg.MailerPass)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	defaultFromEmail = cfg.MailerFromEmail
	defaultToEmail = cfg.MailerToEmail

	return &mailer{dialer: dialer, branch: branch}
}

func (m *mailer) Send(emailToArr []string, subject string, body string) {
	var emailTo string
	if m.branch != "main" {
		emailTo = strings.Join(emailToArr, ",")
	} else {
		emailTo = defaultToEmail
	}

	message := gomail.NewMessage()
	message.SetHeader("From", defaultFromEmail)
	message.SetHeader("To", emailTo)
	message.SetHeader("Subject", subject)
	message.SetBody("text/plain", body)

	if err := m.dialer.DialAndSend(message); err != nil {
		panic(err)
	}
}
