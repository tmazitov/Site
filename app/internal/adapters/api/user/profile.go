package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Profile godoc
// @Summary      User profile data
// @Description  User profile data
// @Tags         User
// @ID           profile
// @Accept       json
// @Produce      json
// @Param        Authorization  header    string  true  "Insert your access token"  default(Bearer <Add access token here>)
// @Success      200            {string}  string  "Success created"
// @Failure      500            {string}  string  "Internal Server Error"
// @Router       /user/profile [get]
func (h *handler) Profile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Security-Policy", "policy")
	w.Header().Set("X-Frame-Options", "DENY")
	w.Header().Set("X-XSS-Protection", "1; mode=block")

	// Get username from access token
	username, _, err := h.JWTHelper.GetUserByToken(r)
	if err != nil {
		log.Println(fmt.Errorf("invalid token: %s", err))
		fmt.Println("")
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// Get all data by username
	user, err := h.userService.GetUserByUsername(username)
	if err != nil {
		log.Println(errors.New("invalid token"))
		fmt.Println("")
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// Write user data to response body
	data := map[string]interface{}{
		"username": user.Username,
		"register": user.Register,
		"email":    user.Email,
		"role":     user.Role,
	}

	json.NewEncoder(w).Encode(data)
}
