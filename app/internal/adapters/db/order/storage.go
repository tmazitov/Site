package order

import (
	"database/sql"
	"site/internal/domain/order"
)

type orderStorage struct {
	Conn *sql.DB
}

func NewStorage(conn *sql.DB) order.Storage {
	return &orderStorage{Conn: conn}
}
