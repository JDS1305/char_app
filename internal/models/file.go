package models

type File struct {
	ID        int    `json:"id"`
	MessageID int    `json:"message_id"`
	URL       string `json:"url"`
	Type      string `json:"type"`
}
