package email

import (
	"bytes"
	"crypto/tls"
	"io"
	"net/smtp"

	"github.com/RedAFD/mega/internal/config"
)

var DefaultSender = &Sender{
	Username: config.EmailDefaultSenderUsername,
	Password: config.EmailDefaultSenderPassword,
	Host:     config.EmailDefaultSenderHost,
	Port:     config.EmailDefaultSenderPort,
}

type Sender struct {
	Username string
	Password string
	Host     string
	Port     string
}

func (s *Sender) Send(to, subject, body string) (err error) {
	buf := &bytes.Buffer{}
	buf.WriteString(s.Host)
	buf.WriteByte(':')
	buf.WriteString(s.Port)
	addr := buf.String()
	buf.Reset()
	buf.WriteString("To: ")
	buf.WriteString(to)
	buf.WriteString("\r\nFrom: ")
	buf.WriteString(s.Username)
	buf.WriteString("\r\nSubject: ")
	buf.WriteString(subject)
	buf.WriteString("\r\nContent-Type: text/html; charset=UTF-8\r\n\r\n")
	buf.WriteString(body)
	data := buf.Bytes()
	auth := smtp.PlainAuth("", s.Username, s.Password, s.Host)

	var conn *tls.Conn
	if conn, err = tls.Dial("tcp", addr, &tls.Config{InsecureSkipVerify: true}); err != nil {
		return
	}
	var client *smtp.Client
	client, err = smtp.NewClient(conn, s.Host)
	if err != nil {
		return
	}
	defer client.Close()
	if ok, _ := client.Extension("AUTH"); ok {
		err = client.Auth(auth)
		if err != nil {
			return
		}
	}
	err = client.Mail(s.Username)
	if err != nil {
		return
	}
	err = client.Rcpt(to)
	if err != nil {
		return
	}
	var w io.WriteCloser
	w, err = client.Data()
	if err != nil {
		return
	}
	_, err = w.Write(data)
	if err != nil {
		w.Close()
		return
	}
	w.Close()
	return client.Quit()
}
