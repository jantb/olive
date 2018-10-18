package rpc

type LanguageChanged struct {
	LanguageID string `json:"language_id"`
	ViewID     string `json:"view_id"`
}
