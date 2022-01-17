package jwt

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"site/internal/domain/models"
	"site/pkg/cache"
	"strings"
	"time"

	"github.com/cristalhq/jwt/v3"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

type UserClains struct {
	jwt.RegisteredClaims
	Username string
	Role     string
}

type RT struct {
	RefreshToken string
}

type helper struct {
	RTCache cache.Repository
}

func NewHelper(RTCache cache.Repository) Helper {
	return &helper{RTCache: RTCache}
}

type Helper interface {
	GenerateAccessToken(u *models.User) (map[string]string, error)
	UpdateRefreshToken(rt RT) (map[string]string, error)
	GetUserByToken(r *http.Request) (string, string, error)
}

func (h *helper) UpdateRefreshToken(rt RT) (map[string]string, error) {
	defer h.RTCache.Del([]byte(rt.RefreshToken))

	userBytes, err := h.RTCache.Get([]byte(rt.RefreshToken))
	if err != nil {
		return nil, fmt.Errorf("get userbytes from refresh: %v", err)
	}

	var u models.User
	err = json.Unmarshal(userBytes, &u)
	if err != nil {
		return nil, fmt.Errorf("marshal error : %v", err)
	}

	return h.GenerateAccessToken(&u)
}

func (h *helper) GenerateAccessToken(u *models.User) (map[string]string, error) {
	key := []byte(viper.GetString("jwt_key"))
	signer, err := jwt.NewSignerHS(jwt.HS256, key)
	if err != nil {
		return nil, err
	}

	builder := jwt.NewBuilder(signer)

	claims := UserClains{
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        u.UUID,
			Audience:  []string{"users"},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 1)),
		},
		Username: u.Username,
		Role:     u.Role,
	}
	token, err := builder.Build(claims)
	if err != nil {
		return nil, err
	}

	refreshTokenUuid := uuid.New()
	userBytes, _ := json.Marshal(u)
	err = h.RTCache.Set([]byte(refreshTokenUuid.String()), userBytes, 0)
	if err != nil {
		return nil, err
	}

	jsonBytes := map[string]string{
		"access_token":  token.String(),
		"refresh_token": refreshTokenUuid.String(),
	}
	if err != nil {
		return nil, err
	}

	return jsonBytes, nil
}

func (h *helper) GetUserByToken(r *http.Request) (string, string, error) {

	authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer")
	if len(authHeader) != 2 {
		return "", "", errors.New("malformed token")
	}
	jwtToken := authHeader[1]

	jwtToken = strings.ReplaceAll(jwtToken, " ", "")

	key := viper.GetString("jwt_key")
	verifier, err := jwt.NewVerifierHS(jwt.HS256, []byte(key))
	if err != nil {
		return "", "", fmt.Errorf("fatal new verifier: %s", err)
	}

	token, err := jwt.ParseAndVerifyString(jwtToken, verifier)
	if err != nil {
		return "", "", fmt.Errorf("fatal parse: %s", err)
	}

	raw := token.RawClaims()

	user := UserClains{}

	err = json.Unmarshal(raw, &user)
	if err != nil {
		return "", "", fmt.Errorf("fatal marshal: %s", err)
	}

	return user.Username, user.Role, nil
}
