package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
	"github.com/nguyenbt456/blocklist/util"
)

// Blocked represent for block infomations
type Blocked struct {
	User []string `json:"user"`
}

// BlockUsers will block users by facebook uid
func BlockUsers(c *gin.Context) {
	pageID, _ := c.GetPostForm("page_id")
	blockedIDs, _ := c.GetPostFormArray("ids")

	blockedData := Blocked{
		User: blockedIDs,
	}

	response, err := util.CreatePageRequest(c, "POST", fbHost, pageID+"/blocked", "", blockedData)
	if err != nil {
		log.Println("Error: ", err)
		return
	}

	temp := map[string]interface{}{}
	json.Unmarshal(response, &temp)

	pp.Println(temp)
	c.JSON(http.StatusOK, temp)
}

// GetBlockedUsers returns users has been blocked from page
func GetBlockedUsers(c *gin.Context) {
	pageID, _ := c.GetPostForm("page_id")

	response, err := util.CreatePageRequest(c, "GET", fbHost, pageID+"/blocked", "", nil)
	if err != nil {
		log.Println("Error: ", err)
		return
	}

	temp := map[string]interface{}{}
	json.Unmarshal(response, &temp)

	pp.Println(temp)
	c.JSON(http.StatusOK, temp)
}
