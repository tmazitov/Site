package user

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Upgrade godoc
// @Summary      Upgrade user role
// @Description  Upgrade user role
// @Tags         User
// @ID           upgrade
// @Accept       json
// @Produce      json
// @Param        Authorization  header    string  true  "Insert your access token"  default(Bearer <Add access token here>)
// @Success      200            {string}  string  "Success created"
// @Failure      500            {string}  string  "Internal Server Error"
// @Router       /user/upgrade [get]
func (h *handler) UpgradeRole(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// Get username from access token
	username, _, err := h.JWTHelper.GetUserByToken(r)
	if err != nil {
		log.Println(fmt.Errorf("invalid token: %s", err))
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// Get all data by username
	user, err := h.userService.GetUserByUsername(username)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	if err = h.userService.UpgradeRole(user); err != nil {
		log.Println(fmt.Errorf("fatal update role: %s", err))
		http.Error(w, "Internal Server Error", 500)
		return
	}

	w.WriteHeader(http.StatusOK)
}
