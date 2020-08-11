package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/nguyenbt456/blocklist/controller"
)

// SetUserToken set user_access_token to global
func SetUserToken(c *gin.Context) {
	token := controller.FBToken.GetUserToken()
	c.Set("user_access_token", token)
}

// SetPageToken set page_access_token to global
func SetPageToken(c *gin.Context) {
	token := controller.FBToken.GetPageToken()
	c.Set("page_access_token", token)
}
