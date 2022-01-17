package jwt

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/cristalhq/jwt/v3"
	"github.com/julienschmidt/httprouter"
	"github.com/spf13/viper"
)

func Middleware(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer")
		if len(authHeader) != 2 {
			log.Println(errors.New("malformed token"))
			fmt.Println("")
			http.Error(w, "Internal Server Error", 500)
			return
		}
		jwtToken := strings.ReplaceAll(authHeader[1], " ", "")
		key := viper.GetString("jwt_key")
		verifer, err := jwt.NewVerifierHS(jwt.HS256, []byte(key))
		if err != nil {
			unauthorized(w, err)
			return
		}

		token, err := jwt.ParseAndVerifyString(jwtToken, verifer)
		if err != nil {
			unauthorized(w, err)
			return
		}

		var uc UserClains

		if err = json.Unmarshal(token.RawClaims(), &uc); err != nil {
			unauthorized(w, err)
			return
		}
		if valid := uc.IsValidAt(time.Now()); !valid {
			unauthorized(w, err)
			return
		}

		h(w, r, ps)
	}
}

func unauthorized(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusUnauthorized)
}
