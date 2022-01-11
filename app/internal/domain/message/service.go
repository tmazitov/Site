package message

import (
	"site/internal/domain/models"

	"site/internal/adapters/api/message"
)

type service struct {
	storage Storage
}

func NewService(stor Storage) message.Service {
	s := &service{storage: stor}
	return s
}

func (s *service) Create(message *models.Message) error {
	return s.storage.Create(message)
}

func (s *service) Change(messageId string) error {
	return s.storage.Change(messageId)
}

func (s *service) Delete(messageId string) error {
	return s.storage.Delete(messageId)
}

func (s *service) List() ([]models.Message, error) {
	return s.storage.List()
}
