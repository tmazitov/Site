package order

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"site/internal/domain/models"

	"github.com/julienschmidt/httprouter"
)

type createParams struct {
	Title       string
	HourCount   int
	Status      string
	FromAddress string
	ToAddress   string
	Comment     string
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("Content-Security-Policy", "policy")
	w.Header().Set("X-Frame-Options", "DENY")
	w.Header().Set("X-XSS-Protection", "1; mode=block")

	var params createParams
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
		log.Println(fmt.Errorf("invalid token: %s", err))
		http.Error(w, "Internal Server Error", 500)
		return
	}

	order := models.Order{
		Writer:      username,
		Title:       params.Title,
		Status:      params.Status,
		FromAddress: params.FromAddress,
		ToAddress:   params.ToAddress,
		Comment:     params.Comment,
		HourCount:   params.HourCount,
	}

	if err := h.orderService.Create(&order); err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	log.Println("POST order-add: ", order.FromAddress)

	w.WriteHeader(http.StatusCreated)
}
