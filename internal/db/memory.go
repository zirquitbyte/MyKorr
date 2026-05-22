package db

import (
	"sync"

	"mykorelay/internal/models"
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

// implement message store interface here
func (db *MemoryDB) Save(m models.Message) error {
	db.mu.Lock()
	defer db.mu.Unlock()

    db.messages = append(db.messages, m)
	return nil
}
