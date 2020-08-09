package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nguyenbt456/blocklist/util"
)

// GetListPages return pages which users are admin or owner
func GetListPages(c *gin.Context) {
	response, err := util.CreateUserRequest(c, "GET", fbHost, "me/accounts", "type=page", nil)
	if err != nil {
		log.Println("Error create Facebook request: ", err)
		return
	}

	fbMetaData := util.HandleFBMetaData(response)

	c.JSON(http.StatusOK, fbMetaData.Data)
}

// ChoosePage choose a page
func ChoosePage(c *gin.Context) {
	pageID, _ := c.GetPostForm("page_id")
	pageToken, _ := c.GetPostForm("page_access_token")

	FBToken.SetPageToken(pageToken)

	http.Redirect(c.Writer, c.Request, "/pages/"+pageID, http.StatusTemporaryRedirect)
}
