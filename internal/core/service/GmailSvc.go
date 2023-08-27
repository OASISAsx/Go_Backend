package service

import (
	config "collection/infrastructure"
	"collection/internal/core/domain"
	"fmt"
	"net/smtp"
)

type sender struct{}

func NewSenderSvc() domain.SenderSvc {
	return sender{}
}

func (s sender) SendEmail(req domain.SenderEmail) error {
	auth := smtp.PlainAuth("", config.GmailUser(), config.GmailPassword(), config.SmtpHost())

	cust := domain.SenderEmail{
		Sender:  config.SmtpFrom(),
		SubJect: req.SubJect,
		Body:    req.Body,
		To:      req.To,
		Image:   req.Image,
	}

	body := fmt.Sprintf(`<p>จาก <b>นักสะสม</b>: คุณได้ทำ %s</p>`, cust.Body)
	body += fmt.Sprintf(`<img src="%s" width="100" height="100" />`, cust.Image)

	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: %s\r\n", cust.Sender)
	msg += fmt.Sprintf("To: %s\r\n", cust.To)
	msg += fmt.Sprintf("Subject: %s\r\n", cust.SubJect)
	msg += fmt.Sprintf("\r\n%s\r\n", body)

	err := smtp.SendMail(config.SmtpAddr(), auth, config.SmtpFrom(), []string{cust.To}, []byte(msg))
	if err != nil {
		fmt.Println(err)
	}
	return err
}
