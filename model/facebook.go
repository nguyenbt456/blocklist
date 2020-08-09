package model

// FBMetaData represent Facebook meta_data
type FBMetaData struct {
	Data   []FBPage `json:"data"`
	Paging FBPaging `json:"paging"`
}
