package user

import (
	"database/sql"
	"fmt"
	"site/internal/domain/models"
	"site/settings"
)

// GetAll get all from users table from the "limit" to the "offset"
func (bs *userStorage) GetAll(offset, limit, timestamp int) ([]*models.User, error) {

	// Select from db
	var rows *sql.Rows
	var err error
	if timestamp != 0 {
		rows, err = settings.DB.Query(`select 
		username, role, register 
		from users
		where "register" between $3 and $4 
		offset $1 limit $2
		`,
			offset, limit, timestamp, timestamp+86400)
	} else {
		rows, err = settings.DB.Query(`select 
		username, role, register 
		from users 
		offset $1 limit $2
		`,
			offset, limit)
	}

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
		err := rows.Scan(&u.Username, &u.Role, &u.Register)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, &u)
	}

	return users, err
}
