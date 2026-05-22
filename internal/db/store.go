package db

import "mykorelay/internal/models"

// interface to be layer between relay and database backends to avoid tight coupling
type MessageStore interface {
	Save(models.Message) error
}
