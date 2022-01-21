package order

import (
	"database/sql"
	"fmt"
	"log"
	"site/internal/domain/models"
)

func (us *orderStorage) List(params map[string]string) ([]*models.Order, error) {

	var rows *sql.Rows
	var err error

	if params["mode"] == "all" {
		rows, err = us.Conn.Query("select uuid, title, writer, worker, date, status, price from orders where worker is null")
		if err != nil {
			e := fmt.Errorf("fatal get list of orders from db: %s", err)
			return nil, e
		}
	} else if params["mode"] == "my" && params["role"] == "Worker" {
		rows, err = us.Conn.Query("select uuid, title, writer, worker, date, status, price from orders where worker=$1 and status != 'done'", params["username"])
		if err != nil {
			e := fmt.Errorf("fatal get list of orders from db: %s", err)
			return nil, e
		}
	} else if params["mode"] == "my" && params["role"] == "Manager" {
		rows, err = us.Conn.Query("select uuid, title, writer, worker, date, status, price from orders where writer=$1 and status != 'done'", params["username"])
		if err != nil {
			e := fmt.Errorf("fatal get list of orders from db: %s", err)
			return nil, e
		}
	}

	// Slice of users
	orders := []*models.Order{}

	// For each row ( user )
	for rows.Next() {
		var worker interface{}
		o := models.Order{}
		// Reading from row user data and writing to u
		err := rows.Scan(
			&o.UUID,
			&o.Title,
			&o.Writer,
			&worker,
			&o.Date,
			&o.Status,
			&o.Price,
		)
		if err != nil {
			log.Println(err)
			continue
		}
		o.Worker = fmt.Sprintf("%v", worker)
		if o.Worker == "<nil>" {
			o.Worker = ""
		}
		orders = append(orders, &o)
	}

	return orders, nil
}
