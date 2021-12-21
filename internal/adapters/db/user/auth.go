package user

import (
	"fmt"
	"site/internal/domain/models"
	"site/settings"
	"time"

	"github.com/google/uuid"
)

func (bs *userStorage) Register(username, password, email string) (*models.User, error) {
	register := time.Now().Unix()

	rand, err := uuid.NewUUID()
	if err != nil {
		return &models.User{}, err
	}
	id := rand.String()

	// Record new user
	_, err = settings.DB.Exec("insert into users (username, password, email, role, register, random) values ($1, $2, $3, $4, $5, $6)",
		username, password, email, "User", register, id)

	if err != nil {
		e := fmt.Errorf("error new user to db: %s", err)
		return &models.User{}, e
	}

	user := models.User{
		Username: username,
		Register: int(register),
		Role:     "User",
	}

	return &user, nil
}

func (bs *userStorage) Login(username, password string) (*models.User, error) {

	rows, err := settings.DB.Query("select password, role, register from users where username=$1", username)

	if err != nil {
		e := fmt.Errorf("error find pass in db: %s", err)
		return &models.User{}, e
	}

	user := models.User{Username: username}
	var pass string
	// For each row
	for rows.Next() {
		// Reading from row user data and writing to u
		err := rows.Scan(&pass, &user.Role, &user.Register)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if pass == password {
			return &user, nil
		}
	}

	return &models.User{}, nil
}
