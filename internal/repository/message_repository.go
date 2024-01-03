package repository

import (
	"database/sql"

	"github.com/greeflas/go_di_example/internal/model"
)

type MessageRepository struct {
	conn *sql.DB
}

func NewMessageRepository(conn *sql.DB) *MessageRepository {
	return &MessageRepository{conn: conn}
}

func (r *MessageRepository) GetMessage() (*model.Message, error) {
	row := r.conn.QueryRow("SELECT id, pattern FROM messages ORDER BY id DESC LIMIT 1")

	message := new(model.Message)
	if err := row.Scan(&message.Id, &message.Pattern); err != nil {
		return nil, err
	}

	return message, nil
}

func (r *MessageRepository) Create(message *model.Message) error {
	_, err := r.conn.Exec("INSERT INTO messages (pattern) VALUES ($1)", message.Pattern)

	return err
}
