package user

import "site/app/internal/domain/models"

type Storage interface {
	Login(username string) (string, error)
	Register(username, password string) error
	GetAll(limit, offset int) ([]*models.User, error)
}
