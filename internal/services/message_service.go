package services

import (
	"chat_app/internal/models"
	"chat_app/internal/repositories"
)

type MessageService struct {
	MessageRepo *repositories.MessageRepository
}

func NewMessageService(messageRepo *repositories.MessageRepository) *MessageService {
	return &MessageService{MessageRepo: messageRepo}
}

func (ms *MessageService) CreateMessage(message *models.Message) error {
	err := ms.MessageRepo.CreateMessage(message)
	if err != nil {
		return err
	}

	return nil
}

func (ms *MessageService) GetMessageByID(messageID int) (*models.Message, error) {
	message, err := ms.MessageRepo.GetMessageByID(messageID)
	if err != nil {
		return nil, err
	}

	return message, nil
}
