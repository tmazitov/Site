package order

import (
	"fmt"
	"site/internal/domain/models"
)

func (us *orderStorage) Get(UUID string) (*models.Order, error) {
	var order models.Order
	var worker interface{}

	row := us.Conn.QueryRow("select * from orders where uuid=$1", UUID)

	err := row.Scan(
		&order.UUID,
		&order.Title,
		&order.Writer,
		&worker,
		&order.Date,
		&order.Status,
		&order.Comment,
		&order.Price,
	)
	if err != nil {
		e := fmt.Errorf("error get order by uuid from db: %s", err)
		return nil, e
	}
	order.Worker = fmt.Sprintf("%v", worker)
	if order.Worker == "<nil>" {
		order.Worker = ""
	}

	return &order, nil
}
