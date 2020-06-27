package common

import (
	"github.com/spf13/viper"
	"net/smtp"
	"strings"
)

func Mail(to, subject, body, mailType string) error {
	user := viper.GetString("stmp.email")
	password := viper.GetString("stmp.pwd")
	host := viper.GetString("stmp.host")
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var contentType string
	if mailType == "html" {
		contentType = "Content-Type: text/" + mailType + "; charset=UTF-8"
	} else {
		contentType = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + "\r\nSubject: " + subject + " \r\n" + contentType + "\r\n\r\n" + body)
	sendTo := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, sendTo, msg)
	return err
}
