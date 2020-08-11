package model

// FBPaging represent Facebook paging
type FBPaging struct {
	Cursors FBCursor `json:"cursors"`
}

// FBCursor represnet Facebook cursors
type FBCursor struct {
	Before string `json:"before"`
	After  string `json:"after"`
}
