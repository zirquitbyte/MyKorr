package db

import (
    "database/sql"

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
reply_to TEXT
);
`
_, err := s.db.Exec(query)
return err
}

func (s *SQLiteStore) Save(msg models.Message) error {
    query := `
    INSERT INTO messages (source, body, timestamp)
    VALUES (?, ?, ?)
    `

    _, err := s.db.Exec(
        query,
        msg.Source,
        msg.Body,
        msg.Timestamp,
    )

    return err
}

func (s *SQLiteStore) GetMessages(string) ([]models.Message, error) {

query := `
SELECT source, body, timestamp
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
    if err := rows.Scan(&msg.Source, &msg.Body, &msg.Timestamp); err != nil {
        return nil, err
    }
    messages = append(messages, msg)
}
return messages, nil

}

func (s *SQLiteStore) Close() error {
    return s.db.Close()
}
