package user

import "site/internal/domain/models"

type Service interface {
	Login(username, password string) (*models.User, error)
	Register(username, password, email string) (*models.User, error)
	GetAll(limit, offset, timestamp int) ([]*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	CheckUsername(username string) error
	CheckEmail(email string) error
	UpgradeRole(user *models.User) error
}
