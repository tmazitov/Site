package user

import (
	"database/sql"
	"site/internal/domain/user"
)

type userStorage struct {
	Conn *sql.DB
}

func NewStorage(conn *sql.DB) user.Storage {
	return &userStorage{Conn: conn}
}
