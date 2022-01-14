package order

import (
	"fmt"
	"log"
	"site/internal/domain/models"
)

func (us *orderStorage) List() ([]*models.Order, error) {

	rows, err := us.Conn.Query("select uuid, title, writer, date, status from orders")
	if err != nil {
		e := fmt.Errorf("fatal get list of orders from db: %s", err)
		return nil, e
	}

	// Slice of users
	orders := []*models.Order{}

	// For each row ( user )
	for rows.Next() {
		o := models.Order{}
		// Reading from row user data and writing to u
		err := rows.Scan(
			&o.UUID,
			&o.Title,
			&o.Writer,
			&o.Date,
			&o.Status,
		)
		if err != nil {
			log.Println(err)
			continue
		}
		orders = append(orders, &o)
	}

	return orders, nil
}
