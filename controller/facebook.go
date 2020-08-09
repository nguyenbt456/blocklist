package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nguyenbt456/blocklist/model"
	"github.com/nguyenbt456/blocklist/util"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
)

// Facebook app info
const (
	fbAppID        = "707994926424988"
	fbAppSecret    = "63f06cba497e4259444448ffd1696dfa"
	fbCallbackURL  = "http://localhost:8000/auth/login/callback"
	fbState        = "helloworld"
	fbResponseType = "code"
	fbHost         = "https://graph.facebook.com/"
)

var fbScopes = []string{
	"public_profile",
	"pages_show_list",
	"pages_read_engagement",
	"pages_read_user_content",
	"pages_manage_engagement",
	"pages_manage_posts",
	"pages_manage_metadata",
}

var oauthCfg = &oauth2.Config{
	ClientID:     util.EVString("FACEBOOK_APP_ID", fbAppID),
	ClientSecret: util.EVString("FACEBOOK_APP_SECRET", fbAppSecret),
	RedirectURL:  fbCallbackURL,
	Scopes:       fbScopes,
	Endpoint:     facebook.Endpoint,
}

// FBToken contains Facebook tokens
var FBToken = &model.FacebookToken{}

// LoginFacebook allow user login to facebook
func LoginFacebook(c *gin.Context) {
	urlString, err := url.Parse(oauthCfg.Endpoint.AuthURL)
	if err != nil {
		log.Println("Error parse: ", err)
		return
	}

	params := url.Values{}
	params.Add("client_id", oauthCfg.ClientID)
	params.Add("redirect_uri", oauthCfg.RedirectURL)
	params.Add("scope", strings.Join(oauthCfg.Scopes, ","))
	params.Add("response_type", fbResponseType)
	params.Add("state", fbState)

	urlString.RawQuery = params.Encode()
	urlResponse := urlString.String()
	log.Println(urlResponse)

	http.Redirect(c.Writer, c.Request, urlResponse, http.StatusTemporaryRedirect)
}

// LoginFacebookCallback ...
func LoginFacebookCallback(c *gin.Context) {
	state := c.Query("state")
	if state != fbState {
		log.Println("oath invalid state:")
		http.Redirect(c.Writer, c.Request, "/", http.StatusTemporaryRedirect)
		return
	}

	code := c.Query("code")
	log.Println("code: ", code)

	token, _ := oauthCfg.Exchange(oauth2.NoContext, code)
	log.Println("token: ", token.AccessToken)
	FBToken.SetUserToken(token.AccessToken)

	response, err := http.Get("https://graph.facebook.com/me?access_token=" + url.QueryEscape(token.AccessToken))
	if err != nil {
		log.Println("Error get accesstoken:", err)
		http.Redirect(c.Writer, c.Request, "/", http.StatusTemporaryRedirect)
		return
	}

	temp := map[string]interface{}{}
	resBody, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(resBody, &temp)
	log.Println("Response: ", temp)

	http.Redirect(c.Writer, c.Request, "/user/pages", http.StatusTemporaryRedirect)
}
