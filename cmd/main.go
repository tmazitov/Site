package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"site/pkg/cache/freecache"
	"site/pkg/middleware/jwt"
	"site/settings"

	user_handler "site/internal/adapters/api/user"
	user_storage "site/internal/adapters/db/user"
	user_service "site/internal/domain/user"

	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
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

	storage := user_storage.NewStorage()
	service := user_service.NewService(storage)

	cache := freecache.NewCacheRepo(104857600)
	helper := jwt.NewHelper(cache)

	handler := user_handler.NewHandler(service, helper)
	handler.Register(router)

	h := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodOptions},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
		//Debug:            true,
	}).Handler(router)

	// Run server
	logrus.Fatal(http.ListenAndServe("localhost:8000", h))
}