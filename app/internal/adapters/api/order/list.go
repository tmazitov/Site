package order

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (h *handler) List(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// Get username from access token
	username, role, err := h.JWTHelper.GetUserByToken(r)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	if role != "Manager" && role != "Admin" && role != "Worker" {
		log.Println(fmt.Errorf("fatal attempt to get list of order without permission by %s", username))
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	mode := r.FormValue("mode")

	p := map[string]string{
		"mode":     mode,
		"username": username,
		"role":     role,
	}

	orders, err := h.orderService.List(p)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	data := map[string]interface{}{
		"orders": orders,
	}

	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Println("fatal encode in userlist: ", err)
		http.Error(w, "Internal Server Error", 500)
	}

}
