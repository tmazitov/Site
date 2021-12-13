package user

import (
	"fmt"
	"math/rand"
	"site/app/settings"
	"time"
)

func (bs *userStorage) Register(username, password string) error {
	register := time.Now().Format(time.RFC1123)

	rand.Seed(time.Now().UnixNano())

	// Record new user
	_, err := settings.DB.Exec("insert into users (username, password, register, random) values ($1, $2, $3, $4)",
		username, password, register, rand.Intn(1000000))

	if err != nil {
		e := fmt.Errorf("error new user to db: %s", err)
		return e
	}

	return nil
}

func (bs *userStorage) Login(username, password string) (bool, error) {

	rows, err := settings.DB.Query("select register from users where username=$1", username)

	if err != nil {
		e := fmt.Errorf("error find pass in db: %s", err)
		return true, e
	}

	var register string
	// For each row
	for rows.Next() {
		// Reading from row user data and writing to u
		err := rows.Scan(&register)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if register != "" {
			return true, nil
		}
	}
	return false, nil
}

func (bs *userStorage) IsExists(username string) (bool, error) {

	rows, err := settings.DB.Query("select random from users where username=$1", username)
	if err != nil {
		return false, err
	}

	var random string
	for rows.Next() {
		// Reading from row user data and writing to u
		err := rows.Scan(&random)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if random != "" {
			return true, nil
		}
	}

	return false, nil
}
