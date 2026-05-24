package db

import (
    "database/sql"
	"time"

    _ "github.com/mattn/go-sqlite3"
	
	"MyKorr/internal/models"
)

type SQLiteStore struct {
    db *sql.DB
}

func NewSQLiteStore(path string) (*SQLiteStore, error) {
	db, err := sql.Open("sqlite3", "data/messages.db")
    if err != nil {
        return nil, err
    }

    return &SQLiteStore{
        db: db,
    }, nil
}

func (s *SQLiteStore) Init() error {
query := `
CREATE TABLE IF NOT EXISTS messages (
id TEXT PRIMARY KEY,
timestamp INTEGER,
conversation_id TEXT,
source TEXT,
sender TEXT,
body TEXT,
attachments INTEGER,
replyto TEXT
);
`
_, err := s.db.Exec(query)
return err
}

func (s *SQLiteStore) Save(msg models.Message) error {
    query := `
    INSERT INTO messages (id, timestamp, conversation_id, source, sender, body, attachments, replyto)
    VALUES (?, ?, ?, ?, ?, ?, ?, ?)
    `

	_, err := s.db.Exec(
	query,
	msg.ID,
	msg.Timestamp.Unix(),
	msg.ConversationID,
	msg.Source,
	msg.Sender,
	msg.Body,
	bool(msg.Attachments), 
	msg.Replyto,
	)

    return err
}

func (s *SQLiteStore) GetMessages(string) ([]models.Message, error) {

query := `
SELECT id, timestamp, conversation_id, source, sender, body, attachments, replyto
FROM messages
ORDER BY timestamp DESC
`

rows, err := s.db.Query(query)
if err != nil {
    return nil, err
}
defer rows.Close()

var messages []models.Message

for rows.Next() {
    var msg models.Message
    var ts int64
    var attachmentsInt int
    if err := rows.Scan(&msg.ID, &ts, &msg.ConversationID, &msg.Source, &msg.Sender, &msg.Body, &attachmentsInt, &msg.Replyto); err != nil {
        return nil, err
    }
    msg.Timestamp = time.Unix(ts, 0)
    msg.Attachments = attachmentsInt != 0
    messages = append(messages, msg)
}

	return messages, nil
}

func (s *SQLiteStore) Close() error {
    return s.db.Close()
}
