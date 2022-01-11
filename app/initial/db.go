package initial

import (
	"database/sql"
	"fmt"
	"github.com/spf13/viper"
)

func initDbTables(db *sql.DB) error {

	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
		username text NOT NULL UNIQUE,
		password text NOT NULL,
		register integer NOT NULL ,
		random text NOT NULL,
		email text NOT NULL UNIQUE,
		role text NOT NULL)`)

	if err != nil {
		e := fmt.Errorf("database create users table: %v", err)
		return e
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS messages (
		uuid text NOT NULL,
		from_user text NOT NULL REFERENCES users (username),
		to_user text NOT NULL REFERENCES users (username),
		date integer NOT NULL,
		message text NOT NULL,
		is_changed bool NOT NULL)`)

	if err != nil {
		e := fmt.Errorf("database create users table: %v", err)
		return e
	}

	return nil
}

func DbConnect() (*sql.DB, error) {

	login := viper.GetString("db_login")
	password := viper.GetString("db_pass")
	name := viper.GetString("db_name")

	var connStr string

	if viper.GetBool("dev") {
		connStr = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", login, password, name)
	} else {
		connStr = fmt.Sprintf("postgresql://%s:%s@postgres/%s?sslmode=disable", login, password, name)
	}

	// Connect to database

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		e := fmt.Errorf("database connect error: %v", err)
		return nil, e
	}
	// Set db conn to setting for other use
	if err := initDbTables(db); err != nil {
		return nil, err
	}

	return db, nil
}