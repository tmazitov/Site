package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"site/app/settings"
	"site/app/start"

	user_handler "site/app/internal/domain/adapters/api/user"
	user_storage "site/app/internal/domain/adapters/db/user"
	user_service "site/app/internal/domain/user"

	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func initDbConnect() error {

	login := viper.GetString("db_login")
	password := viper.GetString("db_pass")
	name := viper.GetString("db_name")

	// Connect to database
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", login, password, name)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		e := fmt.Errorf("database connect error %v", err)
		return e
	}
	// Set db conn to setting for other use
	settings.DB = db

	return nil
}

func initConfig() error {
	viper.AddConfigPath("../configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func main() {
	if err := initConfig(); err != nil {
		logrus.Fatalf("error init configs: %s", err.Error())
	}

	if err := initDbConnect(); err != nil {
		logrus.Fatalf("error init db connect: %s", err.Error())
	}

	// Router init
	router := httprouter.New()

	start.SetEndpoints(router)

	storage := user_storage.NewStorage()
	service := user_service.NewService(storage)
	handler := user_handler.NewHandler(service)
	handler.Register(router)

	// Run server
	logrus.Fatal(http.ListenAndServe("localhost:8000", router))
}
