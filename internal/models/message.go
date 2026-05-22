package models

import "time"

type Message struct {
    Timestamp time.Time
    Event_id	int64
    Source	string
	Sender	string
	Body	string
	Attachments	bool
	Reply_to	int64
}
