package order

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type compliteParams struct {
	UUID string
}

func (h *handler) Complite(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("Content-Security-Policy", "policy")
	w.Header().Set("X-Frame-Options", "DENY")
	w.Header().Set("X-XSS-Protection", "1; mode=block")

	var params compliteParams
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	if err = json.Unmarshal(body, &params); err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// Get username from access token
	username, _, err := h.JWTHelper.GetUserByToken(r)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	order, err := h.orderService.Get(params.UUID)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	if username != order.Writer {
		log.Println(err.Error())
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	if err := h.orderService.Complite(params.UUID); err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	w.WriteHeader(http.StatusOK)
}
