package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"

	"MyKorr/internal/models"
)

type API struct {
	store Store
}

type Store interface {
	Save(m models.Message) error
	GetMessages(conversationID string) ([]models.Message, error)
}

func NewAPI(store Store) *API {
	return &API{store: store}
}

// GET /messages?conversation_id=123
func (a *API) HandleGetMessages(w http.ResponseWriter, r *http.Request) {
	convID := r.URL.Query().Get("conversation_id")

	messages, err := a.store.GetMessages(convID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}
// POST /message
func (a *API) HandlePostMessage(w http.ResponseWriter, r *http.Request) {
	var msg models.Message

	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	msg.ID = uuid.New().String()
	msg.Timestamp = time.Now()

	if msg.ConversationID == "" {
		http.Error(w, "missing conversation_id", http.StatusBadRequest)
		return
	}

	if err := a.store.Save(msg); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)
}
