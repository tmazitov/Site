package user

import "site/app/internal/domain/models"

type Service interface {
	Login(username, password string) (bool, error)
	Register(username, password string) error
	GetAll(limit, offset int) ([]*models.User, error)
	IsExists(username string) (bool, error)
}
