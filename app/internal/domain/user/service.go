package user

import (
	"site/app/internal/domain/adapters/api/user"
	"site/app/internal/domain/models"
)

type service struct {
	storage Storage
}

func NewService(stor Storage) user.Service {
	s := &service{storage: stor}
	return s
}

func (s *service) Register(username, password string) error {
	return s.storage.Register(username, password)
}

func (s *service) Login(username string) (string, error) {
	return s.storage.Login(username)
}

func (s *service) GetAll(limit, offset int) ([]*models.User, error) {
	return s.storage.GetAll(limit, offset)
}