package user

import (
	"fmt"
	"site/settings"
)

func (bs *userStorage) CheckUsername(username string) (bool, error) {

	rows, err := settings.DB.Query("select register from users where username=$1", username)
	if err != nil {
		return false, err
	}

	var register int
	for rows.Next() {
		// Reading from row user data and writing to u
		err := rows.Scan(&register)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if register != 0 {
			return true, nil
		}
	}

	return false, nil
}

func (bs *userStorage) CheckEmail(email string) (bool, error) {
	rows, err := settings.DB.Query("select register from users where email=$1", email)
	if err != nil {
		return false, err
	}

	var register int
	for rows.Next() {
		// Reading from row user data and writing to u
		err := rows.Scan(&register)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if register != 0 {
			return true, nil
		}
	}

	return false, nil
}
