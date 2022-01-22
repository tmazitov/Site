package order

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (h *handler) SetWorker(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get username from access token
	username, role, err := h.JWTHelper.GetUserByToken(r)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	if role != "Worker" {
		log.Println(fmt.Errorf("fatal attempt to update order without permission by %s", username))
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	uuid := r.FormValue("order")

	order, err := h.orderService.Get(uuid)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	if order.Worker != "" {
		log.Printf("fatal set worker: order %v has worker %v", order.Title, order.Worker)
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	newData := map[string]string{
		"worker": username,
		"status": "in process",
	}

	if err := h.orderService.Update(uuid, newData); err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
}
