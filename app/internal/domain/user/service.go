package user

import (
	"site/app/internal/adapters/api/user"
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

func (s *service) Login(username, password string) (bool, error) {
	return s.storage.Login(username, password)
}

func (s *service) GetAll(limit, offset int) ([]*models.User, error) {
	return s.storage.GetAll(limit, offset)
}

func (s *service) IsExists(username string) (bool, error) {
	return s.storage.IsExists(username)
}
