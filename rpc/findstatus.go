package rpc

type FindStatus struct {
	Queries Queries `json:"queries"`
	ViewID  string  `json:"view_id"`
}

type Queries []struct {
	CaseSensitive interface{} `json:"case_sensitive"`
	Chars         interface{} `json:"chars"`
	ID            int         `json:"id"`
	IsRegex       interface{} `json:"is_regex"`
	Matches       int         `json:"matches"`
	WholeWords    interface{} `json:"whole_words"`
}
