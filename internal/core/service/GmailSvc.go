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

	body := fmt.Sprintf(`
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Email Notification</title>
		<style>
		.body {
			position: relative;
			width:auto;
			font-family: Arial, sans-serif;
			background-image: url('https://res.cloudinary.com/dvtvwkwig/image/upload/v1713450594/Collector/i0vrdyuxur9hyat3pzxg.png');
			background-size: cover;
			background-repeat: no-repeat;
			background-position: center;
			height: auto;
			text-align: center;
		}

		.container {
          	position: absolute;
 			inset: 0;
			width:auto;
			padding-top: 60px;
			height: auto;
			padding: 20px;
			background-color: rgba(255, 255, 255, 0.8); /* สีพื้นหลังของเนื้อหา */
			border-radius: 10px;
			box-shadow: 0 0 10px rgba(0, 0, 0, 0.3); /* เงาของกล่อง */
			text-align: center;
			
			
		}

		h2 {
			color: #007bff;
		}

		p {
			margin-bottom: 20px;
		}

		img {
			max-width: 100px;
			height: auto;
			margin-top: 20px;
		}
		
	</style>
	</head>
	<body>
	<div class="body">
		<div class="container" > 
		
			<h2 style="color: #007bff;">แจ้งเตือนผู้ใช้งาน</h2>
			<p>เรียนคุณ %s,</p>
			<p>Email : %s,</p>
			<p>ข้อความนี้เป็นการแจ้งเตือนจากระบบ ซื้อขายของสะสม :</p>
			<p style=" background-color: #f0f0f0; padding: 10px; border-radius: 5px; width: auto;"><b>%s</b></p>
			<img src="%s" alt="Logo" style="max-width: 170px; height: auto; margin-top: 20px;">
			<p>ขอบคุณสำหรับการใช้บริการของเรา</p>
			</div>
		
		</div>
	</body>
	</html>
	`, cust.SubJect, cust.To, cust.Body, cust.Image)

	// สร้างข้อความอีเมล
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
