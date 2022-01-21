package order

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (h *handler) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// Get username from access token
	_, role, err := h.JWTHelper.GetUserByToken(r)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	if role != "Admin" && role != "Manager" {
		log.Println(err.Error())
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	uuid := r.FormValue("order")

	if err := h.orderService.Delete(uuid); err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

}
