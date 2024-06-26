package handler

import (
	"net/http"
	"collection/internal/core/domain"
	"collection/internal/core/service"

	"github.com/gin-gonic/gin"
)

type senderHdl struct {
	svc domain.SenderSvc
}

func NewSenderHdl(svc domain.SenderSvc) senderHdl {
	return senderHdl{
		svc: svc,
	}
}

func (s senderHdl) SendMail() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := domain.SenderEmail{}
		err := c.BindJSON(&req)
		if err != nil {
			c.Error(err)
		}
		err = service.NewSenderSvc().SendEmail(req)
		if err != nil {
			c.Error(err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Emails sent successfully!!",
		})
	}
}