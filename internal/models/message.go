package models

import "time"

type Message struct {
	ID	string	`json:"id"`
	Timestamp time.Time `json:"timestamp"`
	ConversationID	string `json:"conversation_id"`
	Source	string `json:"source"`
	Sender	string `json:"sender"`
	Body	string `json:"body"`
	Attachments	bool
	Replyto	string `json:"replyto"`
}

type Conversation struct {
	ConversationID	string `json:"id"`
	Members []string `json:"members"`
}
