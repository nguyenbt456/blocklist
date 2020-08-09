package util

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateUserRequest send a request to Facebook api server with user access token
func CreateUserRequest(c *gin.Context, method, host, endpoint, query string, data interface{}) ([]byte, error) {
	token, _ := c.Get("user_access_token")

	return createFacebookRequest(c, method, host, endpoint, query, token.(string), data)
}

// CreatePageRequest send a request to Facebook api server with page access token
func CreatePageRequest(c *gin.Context, method, host, endpoint, query string, data interface{}) ([]byte, error) {
	token, _ := c.Get("page_access_token")
	log.Println("Page access token: ", token)

	return createFacebookRequest(c, method, host, endpoint, query, token.(string), data)
}

func createFacebookRequest(c *gin.Context, method, host, endpoint, query, token string, data interface{}) ([]byte, error) {
	url := host + endpoint + "?access_token=" + token + "&" + query

	resBody, err := json.Marshal(data)
	if err != nil {
		resBody = nil
	}

	request, err := http.NewRequest(method, url, bytes.NewBuffer(resBody))
	if err != nil {
		return nil, err
	}
	log.Println(url)

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
