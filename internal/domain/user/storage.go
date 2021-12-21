package user

import "site/internal/domain/models"

type Storage interface {
	Login(username, password string) (*models.User, error)
	Register(username, password, email string) (*models.User, error)
	GetAll(limit, offset int) ([]*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	CheckUsername(username string) (bool, error)
	CheckEmail(email string) (bool, error)
}
