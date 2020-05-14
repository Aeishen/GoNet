package main

import (
"bytes"
"log"
"net/smtp"
)

func main() {
	// Dial 返回一个用于连接到 SMTP 服务器的客户端
	cli, err := smtp.Dial("mail.example.com:25")
	if err != nil {
		log.Fatal(err)
	}

	// 设置 Mail（= 寄件人） 和 Rcpt （= 收件人）
	_ = cli.Mail("sender@example.org")
	_ = cli.Rcpt("recipient@example.net")

	// Data() 返回一个可以写入数据的 writer，这里用 buf.WriteTo(wc) 写入
	wc, err := cli.Data()
	if err != nil {
		log.Fatal(err)
	}
	defer wc.Close()
	buf := bytes.NewBufferString("\"To: recipient@example.net\\r\\nFrom: sender@example.org\\r\\nSubject: 邮件主题\\r\\nContent-Type: text/plain; charset=UTF-8\\r\\n\\r\\nHello World")
	if _, err = buf.WriteTo(wc); err != nil {
		log.Fatal(err)
	}
}