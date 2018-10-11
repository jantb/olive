package rpc

type Line struct {
	Text   string `json:"text"`
	Cursor []int  `json:"cursor"`
	Styles []int  `json:"styles"`
}

type op struct {
	Op    string `json:"op"`
	N     int    `json:"n"`
	Lines []Line `json:"lines"`
}

type update struct {
	Ops      []op `json:"ops"`
	Pristine bool `json:"pristine"`
}

type Update struct {
	ViewID string `json:"view_id"`
	Update update `json:"update"`
}
