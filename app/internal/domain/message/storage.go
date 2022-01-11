package message

import "site/internal/domain/models"

type Storage interface {
	Create(*models.Message) error
	Delete(messageId string) error
	Change(messageId string) error
	List() ([]models.Message, error)
}
