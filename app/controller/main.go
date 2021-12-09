package controller

import (
	"html/template"
	"log"
	"net/http"
	"site/app/model"

	"github.com/julienschmidt/httprouter"
)

// MainPage handler
func MainPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// HTML file paths
	mainPath := "./public/html/main.gohtml"
	commonPath := "./public/html/common.gohtml"
	tablePath := "./public/html/table.gohtml"
	registerPath := "./public/html/register.gohtml"

	// All users (data for page)
	users := model.GetAllUsers()

	// Create a new template
	ts, err := template.ParseFiles(mainPath, commonPath, tablePath, registerPath)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// Name the data values for the page
	data := map[string]interface{}{
		"users": users,
	}

	// Execute the tmplate with prepared data
	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}
