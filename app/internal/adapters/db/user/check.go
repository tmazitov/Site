package user

import (
	"database/sql"
	"errors"
	"fmt"
	"site/settings"
)

func (bs *userStorage) CheckUsername(username string) error {

	row := settings.DB.QueryRow("select register from users where username=$1", username)

	var register int

	err := row.Scan(&register)
	if err != nil && err != sql.ErrNoRows {
		e := fmt.Errorf("error check user by username in db: %s", err)
		return e
	}

	if register != 0 {
		e := errors.New("error user is exists")
		return e
	}

	return nil
}

func (bs *userStorage) CheckEmail(email string) error {
	row := settings.DB.QueryRow("select register from users where email=$1", email)

	var register int

	err := row.Scan(&register)
	if err != nil && err != sql.ErrNoRows {
		e := fmt.Errorf("error check user by email in db: %s", err)
		return e
	}

	if register != 0 {
		e := errors.New("error user is exists")
		return e
	}

	return nil
}
