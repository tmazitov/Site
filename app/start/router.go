package start

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func SetEndpoints(r *httprouter.Router) {

	// Static files
	r.ServeFiles("/public/*filepath", http.Dir("../../public/"))
	// Main page
	r.GET("/", startPage)
}
