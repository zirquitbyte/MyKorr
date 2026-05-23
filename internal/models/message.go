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
	Reply_to	string `json:"reply_to"`
}

type Conversation struct {
	ConversationID	string `json:"id"`
	Members []string `json:"members"`
}
