package user

import (
	"site/internal/domain/user"
)

type userStorage struct {
}

func NewStorage() user.Storage {
	var u userStorage
	return &u
}
