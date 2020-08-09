package util

import (
	"encoding/json"

	"github.com/nguyenbt456/blocklist/model"
)

// ParseInterfaceToMap return map type of input
func ParseInterfaceToMap(value interface{}) *map[string]interface{} {
	result := &(map[string]interface{}{})

	data, _ := json.Marshal(value)
	json.Unmarshal(data, result)

	return result
}

// HandleFBMetaData return FBMetaData
func HandleFBMetaData(metaData []byte) *model.FBMetaData {
	result := &model.FBMetaData{}

	json.Unmarshal(metaData, result)
	return result
}
