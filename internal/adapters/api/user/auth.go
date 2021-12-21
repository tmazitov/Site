package user

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"site/pkg/middleware/jwt"
	"strings"

	"github.com/julienschmidt/httprouter"
)

// GetMD5Hash convert to md5 hash
func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

// Login user handler
/*
POST:
-> username string - name of user
-> password string - user passwod
	OR
PUT:
-> refresh string(cookie)
*/
func (h *handler) SignIn(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Security-Policy", "policy")
	w.Header().Set("X-Frame-Options", "DENY")
	w.Header().Set("X-XSS-Protection", "1; mode=block")

	var tokens map[string]string

	switch r.Method {
	case http.MethodPut: // If you need to update access token

		// Get refresh token
		cookie := r.Header.Get("Cookie")
		refresh := strings.Split(cookie, "refresh_token=")[1]
		refresh = strings.ReplaceAll(refresh, " ", "")
		rt := jwt.RT{RefreshToken: refresh}

		// Get new access and refresh tokens
		newTokens, err := h.JWTHelper.UpdateRefreshToken(rt)
		if err != nil {
			log.Println(err.Error())
			fmt.Println("")
			http.Error(w, "Internal Server Error", 500)
		}

		tokens = newTokens

	case http.MethodPost:
		// Get username and password
		username := r.FormValue("username")
		password := getMD5Hash(r.FormValue("password"))

		// Logining
		user, err := h.userService.Login(username, password)
		if err != nil {
			log.Println(err.Error())
			fmt.Println("")
			http.Error(w, "Internal Server Error", 500)
			return
		}
		// If invalid passwod or username
		if user.Register == 0 {
			fmt.Println("User not exists")
			http.Error(w, "Internal Server Error", http.StatusUnauthorized)
			return
		}

		// Get new access and refresh tokens
		tokens, err = h.JWTHelper.GenerateAccessToken(user)
		if err != nil {
			log.Println(err.Error())
			fmt.Println("")
			http.Error(w, "Internal Server Error", 500)
			return
		}
	}

	// Set refresh token as cookie
	cookie := &http.Cookie{
		Name:     "refresh_token",
		Value:    tokens["refresh_token"],
		MaxAge:   86400 * 30,
		HttpOnly: true,
		Path:     "/",
		Domain:   ".localhost",
	}
	http.SetCookie(w, cookie)

	// Set access token to response body
	data := map[string]string{
		"access_token": tokens["access_token"],
	}

	json.NewEncoder(w).Encode(data)
}

// Add user handler
/*
-> username string - name of user
-> password string - user passwod
-> email    string - user email
*/
func (h *handler) SignUp(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Security-Policy", "policy")
	w.Header().Set("X-Frame-Options", "DENY")
	w.Header().Set("X-XSS-Protection", "1; mode=block")

	// Get username
	username := r.FormValue("username")

	// Checking the existence of a user with this username
	isExists, err := h.userService.IsExists(username)
	if err != nil {
		log.Println(err.Error())
		fmt.Println("")
		http.Error(w, "Internal Server Error", 500)
		return
	}
	if isExists {
		http.Error(w, "User alredy exists", http.StatusForbidden)
		return
	}

	// Get password and email
	password := getMD5Hash(r.FormValue("password"))
	email := r.FormValue("email")

	fmt.Printf("POST user-add: %s \n", username)

	// Record new user
	user, err := h.userService.Register(username, password, email)
	if err != nil {
		log.Println(err.Error())
		fmt.Println("")
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// Generate new access and refresh tokens
	tokens, err := h.JWTHelper.GenerateAccessToken(user)
	if err != nil {
		log.Println(err.Error())
		fmt.Println("")
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// Set refresh token as cookie
	cookie := &http.Cookie{
		Name:     "refresh_token",
		Value:    tokens["refresh_token"],
		MaxAge:   86400 * 30,
		HttpOnly: true,
		Path:     "/",
		Domain:   ".localhost",
	}
	http.SetCookie(w, cookie)

	// Set access token to response body
	data := map[string]string{
		"access_token": tokens["access_token"],
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}
