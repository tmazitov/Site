package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"site/app/controller/start"
	"site/app/controller/user"
	"site/app/settings"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

func init() {
	// Loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	// Get login and password from .env
	username, exists := os.LookupEnv("LOGIN")
	if !exists {
		return
	}

	password, exists := os.LookupEnv("PASSWORD")
	if !exists {
		return
	}

	database, exists := os.LookupEnv("PASSWORD")
	if !exists {
		return
	}

	// Connect to database
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", username, password, database)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		e := fmt.Errorf("database connect error %v", err)
		panic(e)
	}
	// Set db conn to setting for other use
	settings.DB = db

	if err != nil {
		panic(err)
	}
}

func setEndpoints(r *httprouter.Router) {
	// For static files
	r.ServeFiles("/public/*filepath", http.Dir("public"))
	// GETs
	r.GET("/", start.MainPage)
	// POSTs
	r.POST("/user/new", user.Add)
	r.POST("/user/list", user.List)
}

func main() {
	// Router init
	router := httprouter.New()
	// Set rounter endpoints
	setEndpoints(router)
	// Run server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
