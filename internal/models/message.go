package models

type Message struct {
	ID         int    `json:"id"`
	SenderID   int    `json:"sender_id"`
	GroupID    int    `json:"group_id"`
	Content    string `json:"content"`
	Status     string `json:"status"`
	Attachment string `json:"attachment"`
}
