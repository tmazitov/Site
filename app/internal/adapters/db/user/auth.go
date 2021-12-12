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

	// ??? maybe user in sprintf
	query := fmt.Sprintf("insert into users (username, password, register, random) values ('%s', '%s', '%s', %v)",
		username, password, register, rand.Intn(1000000))

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

func (bs *userStorage) IsExists(username string) (bool, error) {

	rows, err := settings.DB.Query("select random from users where username='" + username + "'")
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
