package start

import (
	"log"
	"net/http"
	"site/app/settings"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

func startPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// HTML file paths
	// ??? GOPATH mayby
	mainPath := "../../public/html/main.gohtml"
	commonPath := "../../public/html/common.gohtml"
	tablePath := "../../public/html/table.gohtml"
	registerPath := "../../public/html/register.gohtml"

	// Create a new template
	ts, err := template.ParseFiles(mainPath, commonPath, tablePath, registerPath)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// Name the data values for the page
	data := map[string]interface{}{
		"consts": map[string]interface{}{
			"USER_ROW_COUNT": settings.USER_ROW_COUNT,
		},
	}

	// Execute the tmplate with prepared data
	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}
