package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"site/pkg/cache/freecache"
	"site/pkg/middleware/jwt"

	user_handler "site/internal/adapters/api/user"
	user_storage "site/internal/adapters/db/user"
	user_service "site/internal/domain/user"

	_ "site/docs"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"github.com/spf13/viper"
)

func initDbConnect() (*sql.DB, error) {

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
		e := fmt.Errorf("database connect error %v", err)
		return nil, e
	}
	// Set db conn to setting for other use

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		username text NOT NULL,
		password text NOT NULL,
		register integer NOT NULL ,
		random text NOT NULL,
		email text NOT NULL,
		role text NOT NULL)`)

	if err != nil {
		return nil, err
	}
	return db, nil
}

func initConfig() error {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

// @title        Go Restful API with Swagger
// @version      1.0
// @description  Simple swagger implementation in Go HTTP

// @contact.name   Linggar Primahastoko
// @contact.url    https://linggar.asia
// @contact.email  x@linggar.asia

// @securitydefinitions.basic

// @host      3117-176-52-104-232.ngrok.io/api
// @BasePath  /
func main() {

	log.Println("Set configs")

	if err := initConfig(); err != nil {
		log.Fatalf("error init configs: %s", err.Error())
	}

	log.Println("Set database connection")

	var conn *sql.DB
	conn, err := initDbConnect()
	if err != nil {
		log.Fatalf("error init db connect: %s", err.Error())
	}

	log.Println("Set http router")

	// Router init
	router := httprouter.New()

	storage := user_storage.NewStorage(conn)
	service := user_service.NewService(storage)

	cache := freecache.NewCacheRepo(104857600)
	helper := jwt.NewHelper(cache)

	handler := user_handler.NewHandler(service, helper)
	handler.Register(router)

	router.GET("/swagger/:any", swaggerHandler)

	log.Println("Set cors policy")

	h := cors.New(cors.Options{
		AllowedOrigins:   []string{viper.GetString("front_domain")},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodOptions},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
		//Debug:            true,
	}).Handler(router)

	log.Printf("Starting server on %v port \n", viper.GetString("port"))

	// Run server
	log.Fatal(http.ListenAndServe(":"+viper.GetString("port"), h))
}

func swaggerHandler(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	httpSwagger.WrapHandler(res, req)
}
