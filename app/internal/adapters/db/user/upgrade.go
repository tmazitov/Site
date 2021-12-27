package user

import (
	"fmt"
	"site/internal/domain/models"
	"site/settings"
)

func (bs *userStorage) UpgradeRole(user *models.User) error {

	var err error

	if user.Role == "User" {
		_, err = settings.DB.Exec("update users set role=$1 where role=$2 username=$3", "Moder", user.Role, user.Username)
	} else if user.Role == "Moder" {
		_, err = settings.DB.Exec("update users set role=$1 where role=$2 username=$3", "Admin", user.Role, user.Username)
	} else {
		return nil
	}

	if err != nil {
		e := fmt.Errorf("error upgrade user role : %s", err)
		return e
	}

	return nil
}
