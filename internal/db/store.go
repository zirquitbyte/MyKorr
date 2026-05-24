package db

import "MyKorr/internal/models"

// interface to be layer between relay and database backends to avoid tight coupling
type MessageStore interface {
	Init() error
	Close() error
	GetMessages(conversationID string) ([]models.Message, error)
	Save(models.Message) error
}
