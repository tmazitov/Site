package user

import (
	"fmt"
	"site/internal/domain/models"
	"site/settings"
)

func (bs *userStorage) GetUserByUsername(username string) (*models.User, error) {

	user := &models.User{Username: username}

	row := settings.DB.QueryRow("select role, email, register from users where username=$1", username)

	err := row.Scan(&user.Role, &user.Email, &user.Register)
	if err != nil {
		e := fmt.Errorf("error get user by username in db: %s", err)
		return nil, e
	}

	return user, nil

}
