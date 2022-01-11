package message

import (
	"fmt"
	"site/internal/domain/models"
	"time"

	"github.com/google/uuid"
)

func (us *messageStorage) Create(message *models.Message) error {
	register := time.Now().Unix()

	rand, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	id := rand.String()

	message.UUID = id
	message.Date = int(register)

	// Record new user
	_, err = us.Conn.Exec("insert into users (username, password, email, role, register, random) values ($1, $2, $3, $4, $5, $6)",
		message)

	if err != nil {
		e := fmt.Errorf("fatal create new message in db: %s", err)
		return e
	}

	return nil
}
