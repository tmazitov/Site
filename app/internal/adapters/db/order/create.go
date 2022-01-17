package order

import (
	"fmt"
	"site/internal/domain/models"
	"time"

	"github.com/google/uuid"
)

func (us *orderStorage) Create(order *models.Order) error {

	// Set default data
	order.Date = time.Now().Unix()
	order.Status = "wait"

	rand, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	order.UUID = rand.String()

	// Record new user
	_, err = us.Conn.Exec("insert into orders (uuid, title, writer, date, status, comment, price) values ($1, $2, $3, $4, $5, $6, $7 )",
		order.UUID, order.Title, order.Writer, order.Date, order.Status, order.Comment, order.Price)

	if err != nil {
		e := fmt.Errorf("fatal create order record in db: %s", err)
		return e
	}

	return nil
}
