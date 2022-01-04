package user

import (
	"fmt"
	"site/internal/domain/models"
)

func (us *userStorage) UpgradeRole(user *models.User) error {

	var err error

	if user.Role == "User" {
		_, err = us.Conn.Exec("update users set role=$1 where role=$2 username=$3", "Moder", user.Role, user.Username)
	} else if user.Role == "Moder" {
		_, err = us.Conn.Exec("update users set role=$1 where role=$2 username=$3", "Admin", user.Role, user.Username)
	} else {
		return nil
	}

	if err != nil {
		e := fmt.Errorf("error upgrade user role : %s", err)
		return e
	}

	return nil
}
