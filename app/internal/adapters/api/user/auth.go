package user

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/julienschmidt/httprouter"
)

// GetMD5Hash convert to md5 hash
func getMD5Hash(text string) string {
	hash := sha256.Sum256([]byte(text))
	return hex.EncodeToString(hash[:])
}

type SignInParams struct {
	Username string `json:"username" example:"example" valid:"type(string)"`
	Password string `json:"password" example:"123456789" valid:"type(string)"`
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
	w.Header().Set("Content-Security-Policy", "policy")
	w.Header().Set("X-Frame-Options", "DENY")
	w.Header().Set("X-XSS-Protection", "1; mode=block")

	var tokens map[string]string

	// Get payload
	var params SignInParams
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 102400))
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

	isValid, err := govalidator.ValidateStruct(params)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	if !isValid {
		log.Println("Invalid params for auth")
		http.Error(w, "Bad request", 400)
		return
	}

	params.Password = getMD5Hash(params.Password)

	// Logining
	user, err := h.userService.Login(params.Username, params.Password)
	if err != nil || user == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Get new access and refresh tokens
	tokens, err = h.JWTHelper.GenerateAccessToken(user)
	if err != nil {
		log.Println(err.Error())
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
	data := map[string]interface{}{
		"access_token": tokens["access_token"],
		"expires_in":   time.Now().Unix() + 60*1,
	}

	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Println("fatal encode in auth: ", err)
		http.Error(w, "Internal Server Error", 500)
	}
}

type SignUpParams struct {
	Username string `json:"username" example:"example" minLength:"5" valid:"type(string)"`
	Email    string `json:"email" example:"example@gmail.com" valid:"email"`
	Password string `json:"password" example:"123456789" minLength:"8" valid:"type(string)"`
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
	w.Header().Set("Content-Security-Policy", "policy")
	w.Header().Set("X-Frame-Options", "DENY")
	w.Header().Set("X-XSS-Protection", "1; mode=block")

	// Get payload
	var params SignUpParams
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 102400))
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

	isValid, err := govalidator.ValidateStruct(params)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	if !isValid {
		log.Println("Invalid params for auth")
		http.Error(w, "Bad request", 400)
		return
	}

	params.Password = getMD5Hash(params.Password)

	// Checking the existence of a user with this username

	if err := h.userService.CheckUsername(params.Username); err != nil {
		log.Println(err.Error())
		http.Error(w, "Bad request", 400)
		return
	}

	// Checking the existence of a user with this email

	if err = h.userService.CheckEmail(params.Email); err != nil {
		log.Println(err.Error())
		http.Error(w, "Bad request", 400)
		return
	}

	log.Printf("POST user-add: %s \n", params.Username)

	// Record new user
	user, err := h.userService.Register(params.Username, params.Password, params.Email)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// Generate new access and refresh tokens
	tokens, err := h.JWTHelper.GenerateAccessToken(user)
	if err != nil {
		log.Println(err.Error())
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
	data := map[string]interface{}{
		"access_token": tokens["access_token"],
		"expires_in":   time.Now().Unix() + 60*1,
	}

	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Println("fatal encode in auth: ", err)
		http.Error(w, "Internal Server Error", 500)
	}
}

func (h *handler) SignOut(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	cookie := &http.Cookie{
		Name:     "refresh_token",
		Value:    "",
		MaxAge:   0,
		HttpOnly: true,
		Path:     "/",
		Domain:   ".localhost",
		Expires:  time.Unix(0, 0),
	}

	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusOK)
}
