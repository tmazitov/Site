package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"site/pkg/middleware/jwt"
	"strings"

	"github.com/julienschmidt/httprouter"
)

// Refresh godoc
// @Summary      Refresh user token
// @Description  Refresh user token
// @Tags         User
// @ID           refresh
// @Accept       json
// @Produce      json
// @Success      200  {string}  string  "Success created"
// @Failure      500  {string}  string  "Internal Server Error"
// @Router       /user/refresh [put]
func (h *handler) Refresh(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Security-Policy", "policy")
	w.Header().Set("X-Frame-Options", "DENY")
	w.Header().Set("X-XSS-Protection", "1; mode=block")

	var tokens map[string]string

	// Get refresh token
	refresh := r.Header.Get("Cookie")
	refresh = strings.Split(refresh, "refresh_token=")[1]
	refresh = strings.ReplaceAll(refresh, " ", "")
	rt := jwt.RT{RefreshToken: refresh}

	// Get new access and refresh tokens
	tokens, err := h.JWTHelper.UpdateRefreshToken(rt)
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

	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Println("fatal encode in refresh token: ", err)
		http.Error(w, "Internal Server Error", 500)
	}
}
