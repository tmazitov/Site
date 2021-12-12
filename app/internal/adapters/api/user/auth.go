package user

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// GetMD5Hash convert to md5 hash
func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func (h *handler) SignIn(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

// Add handler
/*
-> username string - name of user
-> password string - user passwod
-> register string - date of user register
-> randomid int    - random value
*/
func (h *handler) SignUp(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("Content-Type", "application/json")

	// Get values about new user
	username := r.FormValue("username")

	isExists, err := h.userService.IsExists(username)

	if err != nil {
		log.Println(err.Error())
		fmt.Println("")
		http.Error(w, "Internal Server Error", 500)
	}

	if isExists {
		http.Error(w, "User alredy exists", http.StatusForbidden)
		return
	}

	password := getMD5Hash(r.FormValue("password"))

	fmt.Printf("POST user-add: %s \n", username)

	// Record new user
	err = h.userService.Register(username, password)

	if err != nil {
		log.Println(err.Error())
		fmt.Println("")
		http.Error(w, "Internal Server Error", 500)
	}
}
