package user

import (
	"fmt"
	"site/app/internal/domain/models"
	"site/app/settings"
)

// GetAll get all from users table from the "limit" to the "offset"
func (bs *userStorage) GetAll(offset, limit int) ([]*models.User, error) {

	// Select from db
	rows, err := settings.DB.Query("select username, register, random from users offset $1 limit $2", offset, limit)

	if err != nil {
		e := fmt.Errorf("error get rows from db: %s", err)
		return nil, e
	}

	// Slice of users
	users := []*models.User{}

	// For each row ( user )
	for rows.Next() {
		u := models.User{}
		// Reading from row user data and writing to u
		err := rows.Scan(&u.Username, &u.Register, &u.UUID)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, &u)
	}

	return users, err
}
