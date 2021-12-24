package user

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// GetMD5Hash convert to md5 hash
func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

type SignInParams struct {
	Username string `json:"username" example:"example"`
	Password string `json:"password" example:"123456789"`
}

// SignIn godoc
// @Summary      User auth
// @Description  User auth
// @Tags         User
// @ID           sign-in
// @Accept       json
// @Produce      json
// @Param        Params  body      SignInParams  true  "SignIn params"
// @Success      200     {string}  string        "Success"
// @Failure      401     {string}  string        "User not exists"
// @Failure      500     {string}  string        "Internal Server Error"
// @Router       /user/entry [post]
func (h *handler) SignIn(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Security-Policy", "policy")
	w.Header().Set("X-Frame-Options", "DENY")
	w.Header().Set("X-XSS-Protection", "1; mode=block")

	var tokens map[string]string

	// Get payload
	var params SignInParams
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
		fmt.Println("")
		http.Error(w, "Internal Server Error", 500)
		return
	}
	if err = json.Unmarshal(body, &params); err != nil {
		log.Println(err.Error())
		fmt.Println("")
		http.Error(w, "Internal Server Error", 500)
		return
	}
	params.Password = getMD5Hash(params.Password)

	// Logining
	user, err := h.userService.Login(params.Username, params.Password)
	if err != nil {
		log.Println(err.Error())
		fmt.Println("")
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// If invalid passwod or username
	if user.Register == 0 {
		http.Error(w, "User not exists", http.StatusUnauthorized)
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

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

type SignUpParams struct {
	Username string `json:"username" example:"example" minLength:"5"`
	Email    string `json:"email" example:"example@gmail.com"`
	Password string `json:"password" example:"123456789" minLength:"8"`
}

// SignUp godoc
// @Summary      Create new user
// @Description  Create new user
// @Tags         User
// @ID           sign-up
// @Accept       json
// @Produce      json
// @Param        Params  body      SignUpParams  true  "SignUp params"
// @Success      201     {string}  string        "Success created"
// @Failure      403     {string}  string        "This username/email is taken"
// @Failure      500     {string}  string        "Internal Server Error"
// @Router       /user/new [post]
func (h *handler) SignUp(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Security-Policy", "policy")
	w.Header().Set("X-Frame-Options", "DENY")
	w.Header().Set("X-XSS-Protection", "1; mode=block")

	// Get payload
	var params SignUpParams
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
		fmt.Println("")
		http.Error(w, "Internal Server Error", 500)
		return
	}

	if err = json.Unmarshal(body, &params); err != nil {
		log.Println(err.Error())
		fmt.Println("")
		http.Error(w, "Internal Server Error", 500)
		return
	}

	params.Password = getMD5Hash(params.Password)

	// Checking the existence of a user with this username
	isExists, err := h.userService.CheckUsername(params.Username)
	if err != nil {
		log.Println(err.Error())
		fmt.Println("")
		http.Error(w, "Internal Server Error", 500)
		return
	}
	if isExists {
		http.Error(w, "This username is taken", http.StatusForbidden)
		return
	}

	// Checking the existence of a user with this email
	isExists, err = h.userService.CheckEmail(params.Email)
	if err != nil {
		log.Println(err.Error())
		fmt.Println("")
		http.Error(w, "Internal Server Error", 500)
		return
	}
	if isExists {
		http.Error(w, "This email is taken", http.StatusForbidden)
		return
	}

	log.Printf("POST user-add: %s \n", params.Username)

	// Record new user
	user, err := h.userService.Register(params.Username, params.Password, params.Email)
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
