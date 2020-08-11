package app

import (
	"github.com/gin-gonic/gin"
	"github.com/nguyenbt456/blocklist/controller"
	"github.com/nguyenbt456/blocklist/middleware"
)

// InitRouter init router for app
func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", controller.HandleMain)

	authentication := router.Group("/auth")
	{
		authentication.GET("/login", controller.LoginFacebook)
		authentication.GET("/login/callback", controller.LoginFacebookCallback)
	}

	user := router.Group("/user")
	{
		user.Use(middleware.SetUserToken)
		user.GET("/pages", controller.GetListPages)
		user.POST("/pages", controller.ChoosePage)
	}

	pages := router.Group("/pages")
	{
		pages.Use(middleware.SetPageToken)
		pages.GET("/blocked", controller.GetBlockedUsers)
		pages.POST("/blocked", controller.BlockUsers)
	}

	return router
}
