package order

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (h *handler) List(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	orders, err := h.orderService.List()
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
