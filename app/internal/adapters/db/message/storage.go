package message

import (
	"database/sql"
	"site/internal/domain/message"
)

type messageStorage struct {
	Conn *sql.DB
}

func NewStorage(conn *sql.DB) message.Storage {
	return &messageStorage{Conn: conn}
}
