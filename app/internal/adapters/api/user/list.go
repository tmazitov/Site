package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type ListParams struct {
	Timestamp int `json:"timestamp" default:"0" valid:"type(int)"`
	Page      int `json:"page" default:"1" valid:"type(int)"`
	Per       int `json:"per_page" default:"16" valid:"type(int)"`
}

// List godoc
// @Summary      User list
// @Description  User list
// @Tags         User
// @ID           list
// @Accept       json
// @Produce      json
// @Param        Authorization  header    string      true  "Insert your access token"  default(Bearer <Add access token here>)
// @Param        Params         query     ListParams  true  "List params"
// @Success      200            {string}  string      "Success created"
// @Failure      403            {string}  string      "Forbidden"
// @Failure      500            {string}  string      "Internal Server Error"
// @Router       /user/list [get]
func (h *handler) List(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Security-Policy", "policy")
	w.Header().Set("X-Frame-Options", "DENY")
	w.Header().Set("X-XSS-Protection", "1; mode=block")

	// Get role of user
	_, role, err := h.JWTHelper.GetUserByToken(r)
	if err != nil {
		log.Println(fmt.Errorf("invalid token: %s", err))
		fmt.Println("")
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// If not Admin
	if role != "Admin" {
		log.Println(errors.New(" attempt to get user list: no valid token"))
		fmt.Println("")
		http.Error(w, "Internal Server Error", http.StatusForbidden)
		return
	}

	var params ListParams
	if params.Page, err = strconv.Atoi(r.FormValue("page")); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if params.Per, err = strconv.Atoi(r.FormValue("per_page")); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if params.Timestamp, err = strconv.Atoi(r.FormValue("timestamp")); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if params.Page < 1 || params.Per < 1 {
		log.Println("Invalid params for userlist")
		http.Error(w, "Bad request", 400)
		return
	}

	log.Printf("POST user-list: part %v , timestamp %v \n", params.Page, params.Timestamp)

	// How much to skip
	offset := params.Per * (params.Page - 1)

	// Get rows with users data
	users, err := h.userService.GetAll(offset, params.Per, params.Timestamp)
	if err != nil {
		log.Println(err.Error())
		fmt.Println("")
		http.Error(w, "Internal Server Error", 500)
	}

	lastPage := params.Page
	if len(users) == params.Per {
		lastPage += 1
	}

	// Write user data to response body
	data := map[string]interface{}{
		"users": users,
	}

	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Println("fatal encode in userlist: ", err)
		http.Error(w, "Internal Server Error", 500)
	}
}
