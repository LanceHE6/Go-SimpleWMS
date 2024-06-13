package utils

import (
	"Go_simpleWMS/config"
	"Go_simpleWMS/utils/emailTemplate"
	"crypto/tls"
	"fmt"
	"gopkg.in/gomail.v2"
	"os"
	"strconv"
)

// SendVerifyEmail 发送绑定邮箱验证邮件
func SendVerifyEmail(target string, account string, code string) error {
	message := emailTemplate.GetVerifyEmailHTML(target, account, code)

	host := os.Getenv("SMTP_HOST")
	if host == "" {
		host = config.ServerConfig.SMTP.HOST
	}
	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if port == 0 {
		port = config.ServerConfig.SMTP.PORT // 使用 SSL/TLS 端口
	}
	userName := os.Getenv("SMTP_USERNAME")
	if userName == "" {
		userName = config.ServerConfig.SMTP.USERNAME
	}
	password := os.Getenv("SMTP_PASSWORD")
	if password == "" {
		password = config.ServerConfig.SMTP.PASSWORD
	}

	m := gomail.NewMessage()
	m.SetHeader("From", "SimpleWMS"+"<"+userName+">")
	m.SetHeader("To", target)
	m.SetHeader("Subject", "绑定邮箱验证码："+code)
	m.SetBody("text/html", message)

	d := gomail.NewDialer(
		host,
		port,
		userName,
		password,
	)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Failed to send bindEmail:", err)
		return err
	} else {
		fmt.Println("Email sent successfully")
		return nil
	}
}
