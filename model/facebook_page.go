package model

// FBPage represent Facebook page
type FBPage struct {
	ID           string       `json:"id"`
	Name         string       `json:"name"`
	AccessToken  string       `json:"access_token"`
	Category     string       `json:"category"`
	CategoryList []FBCategory `json:"category_list"`
	Tasks        []string     `json:"tasks"`
}
