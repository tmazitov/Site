package user

import (
	"site/app/internal/domain/user"
)

type userStorage struct {
}

func NewStorage() user.Storage {
	var u userStorage
	return &u
}
