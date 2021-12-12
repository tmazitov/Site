package user

import (
	"fmt"
	"math/rand"
	"site/app/internal/domain/models"
	"site/app/internal/domain/user"
	"site/app/settings"
	"time"
)

type userStorage struct {
}

func NewStorage() user.Storage {
	var u userStorage
	return &u
}

// GetAll get all from users table from the "limit" to the "offset"
func (bs *userStorage) GetAll(offset, limit int) ([]*models.User, error) {

	query := fmt.Sprintf("select username, register, random from users offset %v limit %v", offset, limit)
	// Select from db
	rows, err := settings.DB.Query(query)

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
		err := rows.Scan(&u.Username, &u.Register, &u.RandomId)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, &u)
	}

	return users, err
}

func (bs *userStorage) Register(username, password string) error {
	register := time.Now()

	rand.Seed(time.Now().UnixNano())

	reg := register.Format("12.12.2021 01:51:36")

	// ??? maybe user in sprintf
	query := fmt.Sprintf("insert into users (username, password, register, random) values ('%s', '%s', '%s', %v)",
		username, password, reg, rand.Intn(1000000))

	// Record new user
	_, err := settings.DB.Exec(query)

	if err != nil {
		e := fmt.Errorf("error new user to db: %s", err)
		return e
	}

	return nil
}

func (bs *userStorage) Login(username string) (string, error) {

	query := fmt.Sprintf("select password from users where username=%v", username)

	rows, err := settings.DB.Query(query)

	if err != nil {
		e := fmt.Errorf("error find pass in db: %s", err)
		return "", e
	}

	var pass string
	// For each row
	for rows.Next() {
		// Reading from row user data and writing to u
		err := rows.Scan(&pass)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if pass != "" {
			break
		}
	}
	return pass, nil
}
