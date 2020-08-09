package model

// FacebookToken represents tokens of Facebook
type FacebookToken struct {
	userAccessToken string
	pageAccessToken string
}

// SetUserToken set value for user_access_token
func (t *FacebookToken) SetUserToken(token string) {
	t.userAccessToken = token
}

// SetPageToken set value for page_access_token
func (t *FacebookToken) SetPageToken(token string) {
	t.pageAccessToken = token
}

// GetUserToken return user access token
func (t *FacebookToken) GetUserToken() string {
	return t.userAccessToken
}

// GetPageToken return page access token
func (t *FacebookToken) GetPageToken() string {
	return t.pageAccessToken
}
