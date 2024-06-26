package middleware

import (
	"collection/errs"
	"collection/logs"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type mdws struct {
}

func New() Middleware {
	return &mdws{}
}

func (m *mdws) ErrorHandler(c *gin.Context) {
	c.Next()
	errors := c.Errors
	for _, err := range errors {
		switch e := err.Err.(type) {
		case *errs.Errs:
			c.AbortWithStatusJSON(e.HTTPStatusCode, e)
		case error:
			c.AbortWithStatusJSON(http.StatusInternalServerError, e)
		}
		return
	}
}

func (m *mdws) Logger(c *gin.Context) {
	t := time.Now()
	logs.Info(fmt.Sprintf("BEGIN| %v | %s%s", c.Request.Method, c.Request.Host, c.Request.URL))
	logs.Info("Executing the proceed....")
	c.Next()
	latency := time.Since(t)
	status := c.Writer.Status()
	logs.Info(fmt.Sprintf("END | %v | %v | %v | %s%s", c.Request.Method, status, latency, c.Request.Host, c.Request.URL))
}
func (m *mdws) Cors(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Max-Age", "86400")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(200)
	} else {
		c.Next()
	}
}
