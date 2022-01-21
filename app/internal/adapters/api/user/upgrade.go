package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type upgradeParams struct {
	Username string
	Role     string
}

// Upgrade godoc
// @Summary      Upgrade user role
// @Description  Upgrade user role
// @Tags         User
// @ID           upgrade
// @Accept       json
// @Produce      json
// @Param        Params  body      upgradeParams  true  "upgrade params"
// @Param        Authorization  header    string  true  "Insert your access token"  default(Bearer <Add access token here>)
// @Success      200            {string}  string  "Success created"
// @Failure      500            {string}  string  "Internal Server Error"
// @Router       /user/upgrade [post]
func (h *handler) UpgradeRole(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var params upgradeParams
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

	// Get role of user
	username, role, err := h.JWTHelper.GetUserByToken(r)
	if err != nil {
		log.Println(fmt.Errorf("invalid token: %s", err))
		http.Error(w, "Internal Server Error", 500)
		return
	}

	if _, err := h.userService.GetUserByUsername(username); err != nil {
		log.Println(err.Error())
		http.Error(w, "Bad request", 400)
		return
	}

	// If not Admin
	if role != "Admin" {
		log.Println(errors.New("attempt to update user role : no valid token"))
		http.Error(w, "Internal Server Error", http.StatusForbidden)
		return
	}

	if err = h.userService.UpgradeRole(params.Username, params.Role); err != nil {
		log.Println(fmt.Errorf("fatal update role: %s", err))
		http.Error(w, "Internal Server Error", 500)
		return
	}

	log.Printf("POST user-upgrade: %s role was upgraded to %s by %s \n", params.Username, params.Role, username)

}
