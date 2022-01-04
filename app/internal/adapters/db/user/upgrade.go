package user

import (
	"fmt"
	"site/internal/domain/models"
)

func (us *userStorage) UpgradeRole(user *models.User) error {

	var err error

	if user.Role == "User" {
		_, err = us.Conn.Exec("update users set role=$1 where username=$2", "Moder", user.Username)
	} else if user.Role == "Moder" {
		_, err = us.Conn.Exec("update users set role=$1 where username=$2", "Admin", user.Username)
	} else {
		return nil
	}

	if err != nil {
		e := fmt.Errorf("error upgrade user role : %s", err)
		return e
	}

	return nil
}
