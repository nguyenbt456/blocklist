package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const htmlIndex = `<html><body>
Logged in with <a href="/auth/login">facebook</a>
</body></html>
`

// HandleMain ...
func HandleMain(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write([]byte(htmlIndex))
}
