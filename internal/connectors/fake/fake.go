package fake

import (
	"fmt"
	"time"

	"mykorelay/internal/db"
	"mykorelay/internal/models"
)

type FakeConnector struct {}

func NewFakeConnector() *FakeConnector {
    return &FakeConnector{}
}

// interface instead of specific DB
func (c *FakeConnector) Start(store db.MessageStore) {
    for {
        msg := models.Message{
			Timestamp: time.Now(),
			Event_id: 1,
			Source: "fake",
			Sender: "ConcernedFriend",
			Body: "So, uhh, you're saying you slept 4 hours last night?",
			Attachments: true,
			Reply_to: 0,
        }

        store.Save(msg)

		fmt.Println("[" + msg.Timestamp.Format("15:04")+ "]", msg.Sender, ":", msg.Body)

        time.Sleep(2 * time.Second)
    }
}
