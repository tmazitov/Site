package user

import (
	"site/internal/adapters/api/user"
	"site/internal/domain/models"
)

type service struct {
	storage Storage
}

func NewService(stor Storage) user.Service {
	s := &service{storage: stor}
	return s
}

func (s *service) Register(username, password, email string) (*models.User, error) {
	return s.storage.Register(username, password, email)
}

func (s *service) Login(username, password string) (*models.User, error) {
	return s.storage.Login(username, password)
}

func (s *service) GetAll(limit, offset, timestamp int) ([]*models.User, error) {
	return s.storage.GetAll(limit, offset, timestamp)
}

func (s *service) CheckUsername(username string) (bool, error) {
	return s.storage.CheckUsername(username)
}

func (s *service) CheckEmail(email string) (bool, error) {
	return s.storage.CheckEmail(email)
}

func (s *service) GetUserByUsername(username string) (*models.User, error) {
	return s.storage.GetUserByUsername(username)
}

func (s *service) UpgradeRole(user *models.User) error {
	return s.storage.UpgradeRole(user)
}
