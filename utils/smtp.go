package utils

import (
	"crypto/tls"
	"strings"

	"github.com/emersion/go-sasl"
	"github.com/emersion/go-smtp"
	"pentag.kr/api-server/configs"
)


var TLS_CONFIG = &tls.Config{
	InsecureSkipVerify: true,
}

func SendEmail(from string, to []string, subject string, body string) error {
	smtpClient, err := smtp.DialStartTLS(configs.Env.SMTPHost + ":" + configs.Env.SMTPPort, TLS_CONFIG)
	if err != nil {
		return err
	}
	auth := sasl.NewLoginClient(configs.Env.SMTPUser, configs.Env.SMTPPassword)
	err = smtpClient.Auth(auth)
	if err != nil {
		return err
	}
	err = smtpClient.Mail("no-reply@pentag.kr", &smtp.MailOptions{
		Body: smtp.Body8BitMIME,
		UTF8: true,
	})
	if err != nil {
		return err
	}

	// build mail header and body
	header := make(map[string]string)
	header["From"] = from
	header["To"] = strings.Join(to, ", ")
	header["Subject"] = subject
	header["Content-Type"] = "text/html; charset=UTF-8"

	// build mail
	var b strings.Builder
	for k, v := range header {
		b.WriteString(k)
		b.WriteString(": ")
		b.WriteString(v)
		b.WriteString("\r\n")
	}
	b.WriteString("\r\n")
	b.WriteString(body)
	r := strings.NewReader(b.String())

	err = smtpClient.SendMail(from, to, r)
	if err != nil {
		return err
	}
	smtpClient.Quit()
	return nil
}