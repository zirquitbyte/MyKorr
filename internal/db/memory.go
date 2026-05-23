package db

import (
	"sync"

	"MyKorr/internal/models"
)

//in-memory storage
type MemoryDB struct {
	mu	sync.Mutex

    messages []models.Message
}

// constructor
func NewMemoryDB() *MemoryDB {
    return &MemoryDB{
	messages: make([]models.Message, 0),
	}
}

func (db *MemoryDB) GetMessages(conversationID string) ([]models.Message, error) {
    db.mu.Lock()
    defer db.mu.Unlock()

    var out []models.Message

    for _, m := range db.messages {
        if m.ConversationID == conversationID {
            out = append(out, m)
        }
    }

    return out, nil
}

// implement message store interface here
func (db *MemoryDB) Save(m models.Message) error {
	db.mu.Lock()
	defer db.mu.Unlock()

    db.messages = append(db.messages, m)
	return nil
}
