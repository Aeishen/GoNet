//需要权限认证并且有多个收件人，可以使用 SendMail 函数

package main

import (
	"log"
	"net/smtp"
)

func main() {

	// 设置认证信息。
	auth := smtp.PlainAuth(
		"",
		"user@example.com",
		"password",
		"mail.example.com",
	)

	// 连接到服务器, 认证, 设置发件人、收件人、发送的内容, 然后发送邮件。
	err := smtp.SendMail(
		"mail.example.com:25",         // 要连接的服务器
		auth,                               // 认证机制
		"sender@example.org",          // 寄件人地址
		[]string{"recipient@example.net"},   // 发件人地址
		[]byte("This is the email body."),   // 发送的消息
	)
	if err != nil {
		log.Fatal(err)
	}
}
