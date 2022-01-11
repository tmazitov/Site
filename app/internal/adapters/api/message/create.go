package message

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"site/internal/domain/models"

	"github.com/julienschmidt/httprouter"
)

type createParams struct {
	ToUsername string
	Message    string
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Security-Policy", "policy")
	w.Header().Set("X-Frame-Options", "DENY")
	w.Header().Set("X-XSS-Protection", "1; mode=block")

	// Get username from access token
	writer, _, err := h.JWTHelper.GetUserByToken(r)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// Get variables about message from request body
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

	message := models.Message{
		FromUser:  writer,
		ToUser:    params.ToUsername,
		Message:   params.Message,
		IsChanged: false,
	}

	if err := h.userService.Create(&message); err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
