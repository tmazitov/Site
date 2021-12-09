package model

import (
	"fmt"
	"site/app/settings"
)

type User struct {
	Username string
	Register string
	RandomId int
}

// GetAllUsers return slice of Users
func GetAllUsers() []User {

	// Select from db
	rows, err := settings.DB.Query("select username, register, random from users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// Slice of users
	users := []User{}

	// For each row ( user )
	for rows.Next() {
		u := User{}
		// Reading from row user data and writing to u
		err := rows.Scan(&u.Username, &u.Register, &u.RandomId)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, u)
	}
	return users
}

func GetPartUsers(offset int, limit int) []User {

	query := fmt.Sprintf("select username, register, random from users limit %v offset %v", limit, offset)
	// Select from db
	rows, err := settings.DB.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// Slice of users
	users := []User{}

	// For each row ( user )
	for rows.Next() {
		u := User{}
		// Reading from row user data and writing to u
		err := rows.Scan(&u.Username, &u.Register, &u.RandomId)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, u)
	}
	return users
}
