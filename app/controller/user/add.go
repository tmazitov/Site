package user

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"site/app/settings"

	"github.com/julienschmidt/httprouter"
)

// GetMD5Hash convert to md5 hash
func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

// Add handler
/*
-> username string - name of user
-> password string - user passwod
-> register string - date of user register
-> randomid int    - random value
*/
func Add(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// Get values about new user
	username := r.FormValue("username")
	password := GetMD5Hash(r.FormValue("password"))
	register := r.FormValue("register")
	random := r.FormValue("random")

	fmt.Printf("POST user-add: %s , %s, %s \n", username, register, random)

	// Record new user
	_, err := settings.DB.Exec("insert into users (username, password, register, random) values ($1, $2, $3, $4)",
		username, password, register, random)

	if err != nil {
		log.Println(err.Error())
		fmt.Println("")
		http.Error(w, "Internal Server Error", 500)
	}
}
