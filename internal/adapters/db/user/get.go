package user

import (
	"fmt"
	"site/internal/domain/models"
	"site/settings"
)

func (bs *userStorage) GetUserByUsername(username string) (*models.User, error) {

	user := &models.User{Username: username}

	rows, err := settings.DB.Query("select role, email, register from users where username=$1", username)
	if err != nil {
		return user, err
	}

	for rows.Next() {
		// Reading from row user data and writing to u
		err := rows.Scan(&user.Role, &user.Email, &user.Register)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}

	return user, nil

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
