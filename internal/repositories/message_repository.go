package repositories

import (
	"chat_app/internal/models"
	"database/sql"
)

type MessageRepository struct {
	DB *sql.DB
}

func NewMessageRepository(db *sql.DB) *MessageRepository {
	return &MessageRepository{DB: db}
}

func (mr *MessageRepository) CreateMessage(message *models.Message) error {
	_, err := mr.DB.Exec("INSERT INTO messages (sender_id, group_id, content, status, attachment) VALUES ($1, $2, $3, $4, $5)",
		message.SenderID, message.GroupID, message.Content, message.Status, message.Attachment)
	if err != nil {
		return err
	}
	return nil
}

func (mr *MessageRepository) GetMessageByID(messageID int) (*models.Message, error) {
	var message models.Message
	row := mr.DB.QueryRow("SELECT id, sender_id, group_id, content, status, attachment FROM messages WHERE id = $1", messageID)
	if err := row.Scan(&message.ID, &message.SenderID, &message.GroupID, &message.Content, &message.Status, &message.Attachment); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}
	return &message, nil
}
