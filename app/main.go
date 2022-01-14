package main

import (
	"database/sql"
	"log"
	"net/http"
	"site/pkg/cache/freecache"
	"site/pkg/middleware/jwt"
	"site/start"

	user_handler "site/internal/adapters/api/user"
	user_storage "site/internal/adapters/db/user"
	user_service "site/internal/domain/user"

	order_handler "site/internal/adapters/api/order"
	order_storage "site/internal/adapters/db/order"
	order_service "site/internal/domain/order"

	_ "site/docs"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"github.com/spf13/viper"
)

// @title        Go Restful API with Swagger
// @version      1.0
// @description  Simple swagger implementation in Go HTTP

// @contact.name   Linggar Primahastoko
// @contact.url    https://linggar.asia
// @contact.email  x@linggar.asia

// @securitydefinitions.basic

// @host      localhost/api
// @BasePath  /
func main() {

	log.Println("Set configs")

	if err := start.Config(); err != nil {
		log.Fatalf("error init configs: %s", err.Error())
	}

	log.Println("Set database connection")

	var conn *sql.DB
	conn, err := start.DbConnect()
	if err != nil {
		log.Fatalf("error init db: %s", err.Error())
	}

	log.Println("Set http router")

	cache := freecache.NewCacheRepo(104857600)
	helper := jwt.NewHelper(cache)

	// Router init
	router := httprouter.New()

	userStorage := user_storage.NewStorage(conn)
	userService := user_service.NewService(userStorage)
	userhandler := user_handler.NewHandler(userService, helper)
	userhandler.Register(router)

	orderStorage := order_storage.NewStorage(conn)
	orderService := order_service.NewService(orderStorage)
	orderHandler := order_handler.NewHandler(orderService, helper)
	orderHandler.Register(router)

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
